# oauth-google

Minimal Google OAuth helpers: exchange authorization code for tokens and
fetch user info using a typed, generic service.

## Install

```go
import "github.com/aatuh/oauth-google"
```

## Quick start

```go
// 1) Exchange authorization code for tokens
svc := oauth.GetOAuthService(oauth.ProviderGoogle)
input := oauth.NewAuthorizationCodeExchangeInput(
  code, clientID, clientSecret, redirectURI,
)
tok, err := svc.ExchangeAuthorizationCode(input)
if err != nil { /* handle */ }

// 2) Fetch user info into your own struct that satisfies UserInfoOutput
type GoogleUserInfo struct {
  Sub               string `json:"sub"`
  Email             string `json:"email"`
  Error             string `json:"error"`
  ErrorDescription  string `json:"error_description"`
}
func (g GoogleUserInfo) GetError() string            { return g.Error }
func (g GoogleUserInfo) GetErrorDescription() string { return g.ErrorDescription }

uiSvc := oauth.GetUserInfoService[GoogleUserInfo](oauth.ProviderGoogle)
info, err := uiSvc.GetUserInfo(tok.AccessToken)
if err != nil { /* handle */ }
_ = info
```

## Notes

- Endpoints used: token (`https://oauth2.googleapis.com/token`) and
  userinfo (`https://www.googleapis.com/oauth2/v3/userinfo`).
- Non‑200 responses are returned as errors.
- `UserInfoService[T]` requires `T` to implement `UserInfoOutput`.
