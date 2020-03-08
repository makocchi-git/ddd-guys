package firestore

import (
	"context"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cloud.google.com/go/firestore"
	domain "github.com/jupemara/ddd-guys/go/domain/model/user"
)

// Repository に状態を持つ形にしようと思ったが、Client の Close処理 を
// Interface に持たせるべきではないと思ったため、レシーバーそれぞれで初期化を行う形式に修正した
type FirestoreRepository struct{}

func NewFirestoreRepository() *FirestoreRepository {
	return &FirestoreRepository{}
}

func (r *FirestoreRepository) Store(user *domain.User) error {
	existedUser, err := r.FindById(domain.NewId(user.Id()))
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return err
		}
	}
	if existedUser != nil {
		return errors.New("Specified user already exists")
	}
	err = r.upsert(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *FirestoreRepository) FindById(id *domain.Id) (*domain.User, error) {
	ctx, client, err := r.initializeFirestore()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	docsnap, err := client.Doc("users/" + id.Value()).Get(ctx)
	if err != nil {
		return nil, err
	}
	var firestoreUser firestoreUser
	if err := docsnap.DataTo(&firestoreUser); err != nil {
		return nil, err
	}
	return domain.NewUser(id, firestoreUser.FirstName, firestoreUser.LastName)
}

func (r *FirestoreRepository) Update(user *domain.User) error {
	_, err := r.FindById(domain.NewId(user.Id()))
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return errors.New("Specified user is not found")
		}
		return err
	}
	err = r.upsert(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *FirestoreRepository) upsert(user *domain.User) error {
	ctx, client, err := r.initializeFirestore()
	wr, err := client.Collection("users").Doc(user.Id()).Set(ctx, firestoreUser{
		FirstName: user.Name().FirstName(),
		LastName:  user.Name().LastName(),
	})
	if err != nil {
		return err
	}

	fmt.Printf("User: %v %v (ID: %v) is upserted at %v\n",
		user.Name().FirstName(),
		user.Name().LastName(),
		user.Id(),
		wr.UpdateTime)
	return nil
}

func (r *FirestoreRepository) initializeFirestore() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return nil, nil, errors.New("GOOGLE_CLOUD_PROJECT must be set")
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, nil, err
	}
	return ctx, client, nil
}
