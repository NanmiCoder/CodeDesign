package apirequest

type ApiRequest struct {
	AppId       string
	Token       string
	TimeStamp   int64
	OriginalUrl string
}

func (ar ApiRequest) BuildFromUrl(url string) error {
	return nil
}

func (ar ApiRequest) GetOriginalUrl() string {
	return ar.OriginalUrl
}

func (ar ApiRequest) GetAppId() string {
	return ar.AppId
}

func (ar ApiRequest) GetTimeStamp() int64 {
	return ar.TimeStamp
}

func (ar ApiRequest) GetToken() string {
	return ar.Token
}
