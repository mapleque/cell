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

func (this *DataService) Auth(responseType, clientId, redirectUri, scope, state string) (resp map[string]interface{}, err error) {
	resp = map[string]interface{}{}
	resp["code"] = "code"
	resp["state"] = state
	resp["id_token"] = "id_token"
	resp["access_token"] = "2YotnFZFEjr1zCsicMWpAA"
	resp["token_type"] = "bearer"
	resp["expires_in"] = 3600
	resp["refresh_token"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	resp["openid"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	return
}

func (this *DataService) GeneralToken(code string) (resp map[string]interface{}, err error) {
	resp = map[string]interface{}{}
	resp["access_token"] = "2YotnFZFEjr1zCsicMWpAA"
	resp["token_type"] = "bearer"
	resp["expires_in"] = 3600
	resp["refresh_token"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	resp["openid"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	return
}

func (this *DataService) RefreshToken(token string) (resp map[string]interface{}, err error) {
	resp = map[string]interface{}{}
	resp["access_token"] = "2YotnFZFEjr1zCsicMWpAA"
	resp["token_type"] = "bearer"
	resp["expires_in"] = 3600
	resp["refresh_token"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	resp["openid"] = "tGzv3JOkF0XG5Qx2TlKWIA"
	return
}

func (this *DataService) CheckClientToken(token string) bool {
	return true
}

func (this *DataService) GetUserinfoByToken(token string) (resp map[string]interface{}, err error) {
	resp = map[string]interface{}{
		"openid": "openid",
		"email":  "email",
		"name":   "name",
		"avatar": "avatar",
	}
	return
}
