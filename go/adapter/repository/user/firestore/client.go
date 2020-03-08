package firestore

import (
	"context"
	"errors"
	"os"

	"cloud.google.com/go/firestore"
)

func CreateFirestoreClient() (*firestore.Client, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return nil, errors.New("GOOGLE_CLOUD_PROJECT must be set")
	}
	credential := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credential == "" {
		return nil, errors.New("GOOGLE_APPLICATION_CREDENTIALS must be set")
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
