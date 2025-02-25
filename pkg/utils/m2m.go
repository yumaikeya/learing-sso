package utils

import (
	"angya-backend/pkg/constants"
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/carlmjohnson/requests"
)

type (
	res struct {
		IdToken string `json:"id_token"`
	}
)

func GetFosIdToken(ctx context.Context) (idToken *string, err error) {
	if os.Getenv("ENV") == constants.ENV_LOCAL {
		return Ptr("token"), nil
	}

	res := res{}
	if err = requests.URL("/oauth2/token").
		Host(strings.ReplaceAll(os.Getenv("FOS_SSO_URL"), "https://", "")).
		Header("Content-Type", "application/x-www-form-urlencoded").
		Header("Authorization", fmt.Sprintf("Basic %s", Base64Enc(fmt.Sprintf("%s:%s", os.Getenv("FOS_CLIENT_ID"), os.Getenv("FOS_CLIENT_SECRET"))))).
		BodyForm(url.Values{"grant_type": []string{"client_credentials"}, "tenant": []string{"cloud"}}).
		ToJSON(&res).Fetch(ctx); err != nil {
		return nil, err
	}

	return &res.IdToken, nil
}
