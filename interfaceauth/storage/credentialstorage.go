package storage

type CredentialStorage interface {
	GetPasswordByAppId(AppId string) string
}
