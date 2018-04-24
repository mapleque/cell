Cell
====
用户认证授权服务

- [首页](https://cell.mapleque.com)
- [OpenID Configuration](https://cell.mapleque.com/.well-known/openid-configuration)

相关协议
----

- OAuth2.0([RFC6749](https://tools.ietf.org/html/rfc6749))
- JWT([RFC7519](https://tools.ietf.org/html/rfc7519))
- [OpenID Connect 1.0](https://openid.net/specs/openid-connect-discovery-1_0.html)

协议流程：
```
     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+
```

Auth
----

Request:

> GET /auth
> Content-Type: application/x-www-form-urlencoded
> params:
> - response_type
>     - code
>     - token
>     - id_token
>     - code token
>     - code id_token
>     - token id_token
>     - code token id_token
>     - none
> - client_id
> - redirect_uri
> - scope
>     - openid
>     - email
>     - profile
> - state

Code Response:

> 302
> Location redirect_uri with params:
> - code
> - state

Id Token Response:

> 302
> Location redirect_uri with params:
> - id_token
> - state

Token Response:

> 302
> Location redirect_uri with params:
> - access_token
> - token_type
>     - bearer
> - expires_in
> - scope
>     - openid
>     - email
>     - profile
> - state

Error Response:

> 302
> Location redirect_uri with params:
> - error
>     - invalid_request
>     - unauthorized_client
>     - access_denied
>     - unsupported_response_type
>     - invalid_scope
>     - server_error
>     - temporarily_unavailable
> - error_description
> - error_uri
> - state

Token
----

Request:

> POST /token
> Content-Type: application/x-www-form-urlencoded
> Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
> # base64(client_id:client_secret)
> - grant_type
>     - authorization_code
> - code
> - redirect_uri
> - client_id

Response:

> Content-Type: application/json;charset=UTF-8
> Cache-Control: no-store
> Pragma: no-cache
>
> {
>   "access_token":"2YotnFZFEjr1zCsicMWpAA",
>   "token_type":"bearer",
>   "expires_in":3600,
>   "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
>   "example_parameter":"example_value"
> }

Error Response:

> - error
>     - invalid_request
>     - invalid_client
>     - invalid_grant
>     - unauthorized_client
>     - unsupported_grant_type
>     - invalid_scope
> - error_description
> - error_uri

Refresh Token
----
Request:

> POST /token
> Content-Type: application/x-www-form-urlencoded
> Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
> - grant_type
>     - refresh_token
> - refresh_token
> - scope
>     - openid
>     - email
>     - profile

Response:

> Content-Type: application/json;charset=UTF-8
> Cache-Control: no-store
> Pragma: no-cache
>
> {
>   "access_token":"2YotnFZFEjr1zCsicMWpAA",
>   "token_type":"bearer",
>   "expires_in":3600,
>   "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
>   "example_parameter":"example_value"
> }

Error Response:

> - error
>     - invalid_request
>     - invalid_client
>     - invalid_grant
>     - unauthorized_client
>     - unsupported_grant_type
>     - invalid_scope
> - error_description
> - error_uri

Userinfo
----

Request:

> POST /resource
> Content-Type: application/x-www-form-urlencoded
> - access_token

Response:

> Content-Type: application/json;charset=UTF-8
> Cache-Control: no-store
> Pragma: no-cache
>
> {
>   "openid":"openid",
>   "email":"email",
>   "name":"name",
>   "avatar":"avatar"
> }

Error Response:

> - error
>     - invalid_request
>     - invalid_client
>     - invalid_grant
>     - unauthorized_client
> - error_description
> - error_uri
