package storage

type MysqlCredentialStorage struct {
}

func (receiver MysqlCredentialStorage) GetPasswordByAppId(AppId string) string {
	return ""
}
