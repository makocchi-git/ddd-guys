package firestore

type firestoreUser struct {
	FirstName string `firestore:"firstname"`
	LastName  string `firestore:"lastname"`
}
