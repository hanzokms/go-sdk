package test

import (
	"context"
	"testing"

	kms "github.com/hanzokms/go-sdk"
)

func TestBatchCreateSecrets(t *testing.T) {
	client := kms.NewClient(context.Background(), kms.Config{
		SiteUrl:          "http://localhost:8080",
		AutoTokenRefresh: true,
	})

	client.Auth().SetAccessToken("<token>")

	secs, err := client.Secrets().Batch().Create(kms.BatchCreateSecretsOptions{
		Environment: "dev",
		SecretPath:  "/",
		ProjectID:   "06c4d805-ac8a-456f-906f-1319554d15ec",
		Secrets: []kms.BatchCreateSecret{
			{
				SecretKey:   "test3",
				SecretValue: "test1",
			},
			{
				SecretKey:   "test4",
				SecretValue: "test2",
			},
		},
	})

	if err != nil {
		t.Fatalf("Failed to create secrets: %v", err)
	}

	t.Logf("Secrets: %v", secs)
}
