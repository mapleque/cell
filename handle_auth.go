package cell

// GET /auth
// Content-Type: application/x-www-form-urlencoded
// params:
// - response_type
//     - code
//     - token
//     - id_token
//     - code token
//     - code id_token
//     - token id_token
//     - code token id_token
//     - none
// - client_id
// - redirect_uri
// - scope
//     - openid
//     - email
//     - profile
// - state
//
// code response:
// 302
// Location redirect_uri with params:
// - code
// - state
//
// token response:
// 302
// Location redirect_uri with params:
// - access_token
// - token_type
//     - bearer
// - expires_in
// - scope
//     - openid
//     - email
//     - profile
// - state
//
// id_token response:
// - id_token
// - state
//
// error response:
// 302
// Location redirect_uri with params:
// - error
//     - invalid_request
//     - unauthorized_client
//     - access_denied
//     - unsupported_response_type
//     - invalid_scope
//     - server_error
//     - temporarily_unavailable
// - error_description
// - error_uri
// - state
func HandleAuthPage(c *web.Context) {
}

func HandleAuth(c *web.Context) {
}
