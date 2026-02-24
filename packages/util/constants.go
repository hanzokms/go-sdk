package util

import (
	"context"
	"errors"
)

// Auth related:
const (
	KMS_AUTH_ORGANIZATION_SLUG_ENV_NAME = "KMS_AUTH_ORGANIZATION_SLUG"

	// Universal auth:
	KMS_UNIVERSAL_AUTH_CLIENT_ID_ENV_NAME     = "KMS_UNIVERSAL_AUTH_CLIENT_ID"
	KMS_UNIVERSAL_AUTH_CLIENT_SECRET_ENV_NAME = "KMS_UNIVERSAL_AUTH_CLIENT_SECRET"

	// GCP auth:
	KMS_GCP_AUTH_IDENTITY_ID_ENV_NAME                  = "KMS_GCP_AUTH_IDENTITY_ID"
	KMS_GCP_IAM_SERVICE_ACCOUNT_KEY_FILE_PATH_ENV_NAME = "KMS_GCP_IAM_SERVICE_ACCOUNT_KEY_FILE_PATH"

	// AWS auth:
	KMS_AWS_IAM_AUTH_IDENTITY_ID_ENV_NAME = "KMS_AWS_IAM_AUTH_IDENTITY_ID"

	// Azure auth:
	KMS_AZURE_AUTH_IDENTITY_ID_ENV_NAME = "KMS_AZURE_AUTH_IDENTITY_ID"

	// OCI auth:
	KMS_OCI_AUTH_IDENTITY_ID_ENV_NAME = "KMS_OCI_AUTH_IDENTITY_ID"

	// LDAP auth:
	KMS_LDAP_AUTH_IDENTITY_ID_ENV_NAME = "KMS_LDAP_AUTH_IDENTITY_ID"

	// Kubernetes auth:
	KMS_KUBERNETES_IDENTITY_ID_ENV_NAME                = "KMS_KUBERNETES_IDENTITY_ID"
	KMS_KUBERNETES_SERVICE_ACCOUNT_TOKEN_PATH_ENV_NAME = "KMS_KUBERNETES_SERVICE_ACCOUNT_TOKEN_PATH"

	// OIDC auth:
	KMS_OIDC_AUTH_IDENTITY_ID_ENV_NAME = "KMS_OIDC_AUTH_IDENTITY_ID"

	// Access token:
	KMS_ACCESS_TOKEN_ENV_NAME = "KMS_ACCESS_TOKEN"

	// AWS metadata service:
	AWS_EC2_METADATA_TOKEN_URL             = "http://169.254.169.254/latest/api/token"
	AWS_EC2_INSTANCE_IDENTITY_DOCUMENT_URL = "http://169.254.169.254/latest/dynamic/instance-identity/document"

	// Azure metadata service:
	AZURE_METADATA_SERVICE_URL = "http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=" // End of the URL needs to be appended with the resource
	AZURE_DEFAULT_RESOURCE     = "https%3A%2F%2Fmanagement.azure.com/"
)

type AuthMethod string

const (
	ACCESS_TOKEN   AuthMethod = "ACCESS_TOKEN"
	UNIVERSAL_AUTH AuthMethod = "UNIVERSAL_AUTH"
	GCP_ID_TOKEN   AuthMethod = "GCP_ID_TOKEN"
	GCP_IAM        AuthMethod = "GCP_IAM"
	AWS_IAM        AuthMethod = "AWS_IAM"
	KUBERNETES     AuthMethod = "KUBERNETES"
	AZURE          AuthMethod = "AZURE"
	OIDC_AUTH      AuthMethod = "OIDC_AUTH"
	JWT_AUTH       AuthMethod = "JWT_AUTH"
	LDAP_AUTH      AuthMethod = "LDAP_AUTH"
	OCI_AUTH       AuthMethod = "OCI_AUTH"
)

// SSH related:
type CertKeyAlgorithm string

const (
	RSA2048   CertKeyAlgorithm = "RSA_2048"
	RSA4096   CertKeyAlgorithm = "RSA_4096"
	ECDSAP256 CertKeyAlgorithm = "EC_prime256v1"
	ECDSAP384 CertKeyAlgorithm = "EC_secp384r1"
)

type SshCertType string

const (
	UserCert SshCertType = "user"
	HostCert SshCertType = "host"
)

// General:
const (
	DEFAULT_KMS_API_URL                     = "https://kms.hanzo.ai/api"
	DEFAULT_KUBERNETES_SERVICE_ACCOUNT_TOKEN_PATH = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

var ErrContextCanceled = errors.New("context canceled")
var ErrContextDeadlineExceeded error = context.DeadlineExceeded
