package firestore

import (
	ctx "context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/spf13/viper"
)

func CreateClient() *firestore.Client {
	projectID := viper.GetString("PROJECT_ID")

	client, err := firestore.NewClient(ctx.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
