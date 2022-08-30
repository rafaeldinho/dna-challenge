package firestore

import (
	ctx "context"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateClient() *firestore.Client {
	projectID := "control-gastos-296302"

	client, err := firestore.NewClient(ctx.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
