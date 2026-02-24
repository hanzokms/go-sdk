package kms

import (
	api "github.com/hanzokms/go-sdk/packages/api/auth"
	"github.com/hanzokms/go-sdk/packages/errors"
	"github.com/hanzokms/go-sdk/packages/models"
)

type OciAuthLoginOptions struct {
	IdentityID  string
	PrivateKey  string
	Fingerprint string
	UserID      string
	TenancyID   string
	Region      string
	Passphrase  *string
}

type MachineIdentityCredential = api.MachineIdentityAuthLoginResponse

type Secret = models.Secret
type SecretImport = models.SecretImport

type APIError = errors.APIError
type RequestError = errors.RequestError
