package test

// import (
// 	"testing"

// 	kms "github.com/hanzokms/go-sdk"
// )

// func TestKubernetesAuthLogin(t *testing.T) {

// 	t.Skip("Skipping Kubernetes Auth test -- requires running in a Kubernetes cluster")

// 	client, err := kms.NewClient(kms.Config{
// 		SiteUrl: "http://localhost:8080",
// 		Auth: kms.Authentication{
// 			Kubernetes: kms.KubernetesAuth{
// 				IdentityID:              "K8_MACHINE_IDENTITY_ID",
// 				ServiceAccountTokenPath: "/var/run/secrets/kubernetes.io/serviceaccount/token", // Optional
// 			},
// 		},
// 	})

// 	// token1, err := client1.auth.kubernetesLogin(...)
// 	// token2, err := client2.auth.kubernetesLogin(...)

// 	// handle err

// 	// client.auth.setToken(token)

// 	if err != nil {
// 		t.Fatalf("Failed to create KMS client: %v", err)
// 	}

// 	secrets, err := client.Secrets().List(kms.ListSecretsOptions{
// 		ProjectID:   "PROJECT_ID",
// 		Environment: "ENV_SLUG",
// 	})

// 	if err != nil {
// 		t.Fatalf("Failed to list secrets: %v", err)
// 	}

// 	if len(secrets) == 0 {
// 		t.Fatalf("No secrets found")
// 	}

// 	t.Logf("Secrets: %v", secrets)

// }
