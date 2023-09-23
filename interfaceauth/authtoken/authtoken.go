package authtoken

type AuthToken struct {
	Token     string
	TimeStamp int64
}

func (at AuthToken) Match(clientAuthToken AuthToken) bool {
	return false
}

func (at AuthToken) IsExpired() bool {
	return true
}

func (at AuthToken) encrypt(originalUrl, password, appId string) string {
	return ""
}

func Generate(originalUrl string, timeStamp int64, appId string, password string) *AuthToken {
	newAuthToken := &AuthToken{
		TimeStamp: timeStamp,
	}
	encryptedToken := newAuthToken.encrypt(originalUrl, password, appId)
	newAuthToken.Token = encryptedToken

	return newAuthToken
}
