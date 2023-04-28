package gcs

import (
	"context"
	"log"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCS struct {
	*storage.BucketHandle
}

func NewGCS(ctx context.Context) GCS {
	root, _ := os.Getwd()
	filepath := path.Join(root, "credentials.json")

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(filepath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucketName := "kidum-storage"

	bucket := client.Bucket(bucketName)
	return GCS{
		BucketHandle: bucket,
	}
}

func (g *GCS) Upload() {}
