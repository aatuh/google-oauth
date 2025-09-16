package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const userInfoEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

type GoogleUserInfoService[T UserInfoOutput] struct {
	UserInfoEndpoint string
}

func NewGoogleUserInfoService[T UserInfoOutput]() *GoogleUserInfoService[T] {
	return &GoogleUserInfoService[T]{
		UserInfoEndpoint: userInfoEndpoint,
	}
}

func (g *GoogleUserInfoService[T]) GetUserInfo(accessToken string) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, g.UserInfoEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo T
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
