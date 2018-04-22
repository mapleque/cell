package cell

// POST /token
// Content-Type: application/x-www-form-urlencoded
// Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
// # base64(client_id:client_secret)
// grant token params:
// - grant_type
//     - authorization_code
// TODO
// - code
// - redirect_uri
// - client_id
//
// refresh token params:
// - grant_type
//     - refresh_token
// - refresh_token
// - scope
//     - openid
//     - email
//     - profile
//
// response:
// Content-Type: application/json;charset=UTF-8
// Cache-Control: no-store
// Pragma: no-cache
//
// {
//   "access_token":"2YotnFZFEjr1zCsicMWpAA",
//   "token_type":"bearer",
//   "expires_in":3600,
//   "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
//   "example_parameter":"example_value"
// }
// error response:
// - error
//     - invalid_request
//     - invalid_client
//     - invalid_grant
//     - unauthorized_client
//     - unsupported_grant_type
//     - invalid_scope
// - error_description
// - error_uri
func HandleToken(c *web.Context) {
}
