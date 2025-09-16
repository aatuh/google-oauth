package oauth

import "net/url"

type GrantType string

func (g GrantType) String() string {
	return string(g)
}

const GrantTypeAuthorizationCode GrantType = "authorization_code"

type OAuthTokenExchanger interface {
	ExchangeAuthorizationCode(
		input *CodeExchangeInput,
	) (*CodeExchangeOutput, error)
}

type UserInfoGetter[T UserInfoOutput] interface {
	GetUserInfo(accessToken string) (*T, error)
}

type UserInfoOutput interface {
	GetError() string
	GetErrorDescription() string
}

type CodeExchangeInput struct {
	Code         string    `json:"code"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	RedirectURI  string    `json:"redirect_uri"`
	GrantType    GrantType `json:"grant_type"`
}

func NewAuthorizationCodeExchangeInput(
	code string,
	clientID string,
	clientSecret string,
	redirectURI string,
) *CodeExchangeInput {
	return &CodeExchangeInput{
		Code:         code,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		GrantType:    GrantTypeAuthorizationCode,
	}
}

func (c *CodeExchangeInput) ToURLValues() url.Values {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.ClientSecret)
	data.Set("code", c.Code)
	data.Set("grant_type", c.GrantType.String())
	data.Set("redirect_uri", c.RedirectURI)

	return data
}

type CodeExchangeOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	Error        string `json:" error"`
}
