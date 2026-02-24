# Hanzo KMS Go SDK

Official Go SDK for [Hanzo KMS](https://kms.hanzo.ai) -- secret management, encryption, SSH certificates, and dynamic secrets.

## Installation

```bash
go get github.com/hanzokms/go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	kms "github.com/hanzokms/go-sdk"
)

func main() {
	client := kms.NewClient(context.Background(), kms.Config{
		SiteUrl:          "https://kms.hanzo.ai",
		AutoTokenRefresh: true,
	})

	// Authenticate with Universal Auth
	_, err := client.Auth().UniversalAuthLogin(
		os.Getenv("KMS_UNIVERSAL_AUTH_CLIENT_ID"),
		os.Getenv("KMS_UNIVERSAL_AUTH_CLIENT_SECRET"),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "auth failed: %v\n", err)
		os.Exit(1)
	}

	// List secrets
	secrets, err := client.Secrets().List(kms.ListSecretsOptions{
		ProjectID:   os.Getenv("KMS_PROJECT_ID"),
		Environment: "prod",
		SecretPath:  "/",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "list secrets failed: %v\n", err)
		os.Exit(1)
	}

	for _, s := range secrets {
		fmt.Printf("%s = %s\n", s.SecretKey, s.SecretValue)
	}
}
```

## Authentication Methods

| Method | Function |
|--------|----------|
| Universal Auth | `client.Auth().UniversalAuthLogin(clientID, clientSecret)` |
| Access Token | `client.Auth().SetAccessToken(token)` |
| Kubernetes | `client.Auth().KubernetesAuthLogin(identityID, tokenPath)` |
| AWS IAM | `client.Auth().AwsIamAuthLogin(identityID)` |
| Azure | `client.Auth().AzureAuthLogin(identityID, resource)` |
| GCP ID Token | `client.Auth().GcpIdTokenAuthLogin(identityID)` |
| GCP IAM | `client.Auth().GcpIamAuthLogin(identityID, keyPath)` |
| OIDC | `client.Auth().OidcAuthLogin(identityID, jwt)` |
| JWT | `client.Auth().JwtAuthLogin(identityID, jwt)` |
| LDAP | `client.Auth().LdapAuthLogin(identityID, user, pass)` |
| OCI | `client.Auth().OciAuthLogin(options)` |

## Features

- **Secrets** -- CRUD operations, batch create, import support
- **Folders** -- organize secrets into folder hierarchies
- **Dynamic Secrets** -- lease-based credentials with auto-rotation
- **KMS Encryption** -- encrypt/decrypt data, key management, signing/verification
- **SSH Certificates** -- issue and sign SSH certificates, host management
- **Auto Token Refresh** -- background goroutine handles token renewal
- **Caching** -- optional LRU cache for API responses
- **Retry with Backoff** -- configurable exponential backoff on failures

## Environment Variables

| Variable | Description |
|----------|-------------|
| `KMS_UNIVERSAL_AUTH_CLIENT_ID` | Universal Auth client ID |
| `KMS_UNIVERSAL_AUTH_CLIENT_SECRET` | Universal Auth client secret |
| `KMS_ACCESS_TOKEN` | Direct access token |
| `KMS_AUTH_ORGANIZATION_SLUG` | Scoped organization slug |
| `KMS_KUBERNETES_IDENTITY_ID` | Kubernetes identity ID |
| `KMS_AWS_IAM_AUTH_IDENTITY_ID` | AWS IAM identity ID |
| `KMS_AZURE_AUTH_IDENTITY_ID` | Azure identity ID |
| `KMS_GCP_AUTH_IDENTITY_ID` | GCP identity ID |
| `KMS_OIDC_AUTH_IDENTITY_ID` | OIDC identity ID |
| `KMS_LDAP_AUTH_IDENTITY_ID` | LDAP identity ID |
| `KMS_OCI_AUTH_IDENTITY_ID` | OCI identity ID |

## Configuration

```go
client := kms.NewClient(context.Background(), kms.Config{
	SiteUrl:              "https://kms.hanzo.ai", // default
	AutoTokenRefresh:     true,                   // default
	CacheExpiryInSeconds: 300,                    // 5-minute cache
	LogLevel:             kms.LogLevelDebug,       // verbose logging
	SilentMode:           false,                  // show warnings
	CaCertificate:        pemString,              // custom CA
	RetryRequestsConfig: &kms.RetryRequestsConfig{
		ExponentialBackoff: &kms.ExponentialBackoffStrategy{
			BaseDelay:  1 * time.Second,
			MaxRetries: 5,
			MaxDelay:   30 * time.Second,
		},
	},
})
```

## Security

Report vulnerabilities to **security@hanzo.ai**. Do not file public issues for security concerns.

## License

MIT -- see [LICENSE](LICENSE).
