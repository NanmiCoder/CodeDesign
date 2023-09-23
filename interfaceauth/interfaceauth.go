package interfaceauth

import (
	"log"

	"interfaceauth/apirequest"
	"interfaceauth/authtoken"
	"interfaceauth/storage"
)

type ApiAuthenticator interface {
	Auth(url string)
}

func NewDefaultApiAuthenticator() *DefaultApiAuthenticatorImpl {
	return &DefaultApiAuthenticatorImpl{
		apiRequest:        &apirequest.ApiRequest{},
		credentialStorage: &storage.MysqlCredentialStorage{},
	}
}

type DefaultApiAuthenticatorImpl struct {
	apiRequest        *apirequest.ApiRequest
	credentialStorage storage.CredentialStorage
}

func (da *DefaultApiAuthenticatorImpl) Auth(url string) bool {
	// 解析传递过来 url 中的各项信息
	err := da.apiRequest.BuildFromUrl(url)
	if err != nil {
		log.Fatal("build url failed and err:", err)
		return false
	}
	appId := da.apiRequest.GetAppId()
	token := da.apiRequest.GetToken()
	originalUrl := da.apiRequest.GetOriginalUrl()
	timeStamp := da.apiRequest.GetTimeStamp()

	// 创建 client AuthToken 实例并验证时间是否过期
	clientAuthTokenSrv := authtoken.AuthToken{Token: token, TimeStamp: timeStamp}
	if clientAuthTokenSrv.IsExpired() {
		log.Fatal("token is expired !")
		return false
	}

	// 根据原有参数生成一个服务段的 AuthToken 实例，然后与客户端的匹配
	password := da.credentialStorage.GetPasswordByAppId(appId)
	var srvAuthToken *authtoken.AuthToken
	srvAuthToken = authtoken.Generate(originalUrl, timeStamp, token, password)
	if !srvAuthToken.Match(clientAuthTokenSrv) {
		return false
	}
	return true

}
