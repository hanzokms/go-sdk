package test

// import (
// 	"os"
// 	"testing"

// 	kms "github.com/hanzokms/go-sdk"
// )

// func TestUniversalAuthLogin(t *testing.T) {

// 	client, err := kms.NewClient(kms.Config{
// 		Auth: kms.Authentication{
// 			UniversalAuth: kms.UniversalAuth{
// 				ClientID:     os.Getenv("GO_SDK_TEST_UNIVERSAL_AUTH_CLIENT_ID"),
// 				ClientSecret: os.Getenv("GO_SDK_TEST_UNIVERSAL_AUTH_CLIENT_SECRET"),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		t.Fatalf("Failed to create KMS client: %v", err)
// 	}

// 	secrets, err := client.Secrets().List(kms.ListSecretsOptions{
// 		ProjectID:   os.Getenv("GO_SDK_TEST_PROJECT_ID"),
// 		Environment: "dev",
// 	})

// 	if err != nil {
// 		t.Fatalf("Failed to list secrets: %v", err)
// 	}

// 	if len(secrets) == 0 {
// 		t.Fatalf("No secrets found")
// 	}

// 	t.Logf("Secrets: %v", secrets)

// }
