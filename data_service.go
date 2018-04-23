package cell

type IDataService interface {
	CheckClientAuthAndGet(clientId, redirectUri string) (clientName, clientDesc string, ok bool, err error)
	Auth(responseType, clientId, redirectUri, scope, state string) (resp map[string]string, err error)
}

type Client struct {
	ClientId     string
	ClientSecret string
	Name         string
	Desc         string
	RedirectUri  string
}

type DataService struct {
}

func NewDataService() *DataService {
	return &DataService{}
}

func (this *DataService) CheckClientAuthAndGet(clientId, redirectUri string) (clientName, clientDesc string, ok bool, err error) {
	clientName = "client name"
	clientDesc = "client desc"
	ok = true
	err = nil
	return
}

func (this *DataService) Auth(responseType, clientId, redirectUri, scope, state string) (resp map[string]string, err error) {
	resp = map[string]string{}
	resp["code"] = "code"
	resp["state"] = state
	return
}
