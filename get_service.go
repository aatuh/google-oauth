package oauth

type Provider string

const (
	ProviderGoogle Provider = "Google"
)

func GetOAuthService(provider Provider) OAuthTokenExchanger {
	switch provider {
	case ProviderGoogle:
		return NewGoogleOAuthService()
	default:
		return nil
	}
}

func GetUserInfoService[T UserInfoOutput](provider Provider) UserInfoGetter[T] {
	switch provider {
	case ProviderGoogle:
		return NewGoogleUserInfoService[T]()
	default:
		return nil
	}
}
