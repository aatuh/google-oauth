package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const tokenEndpoint = "https://oauth2.googleapis.com/token"

type GoogleOAuthService struct {
	TokenEndpoint string
}

func NewGoogleOAuthService() *GoogleOAuthService {
	return &GoogleOAuthService{
		TokenEndpoint: tokenEndpoint,
	}
}

func (g *GoogleOAuthService) ExchangeAuthorizationCode(
	input *CodeExchangeInput,
) (*CodeExchangeOutput, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		tokenEndpoint,
		strings.NewReader(input.ToURLValues().Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange token: %s", string(body))
	}

	var output CodeExchangeOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
