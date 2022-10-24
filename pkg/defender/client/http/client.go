package defenderclienthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var ErrNotImplemented = fmt.Errorf("Not Implemented")

type Client struct {
	cfg *Config

	csrp    *cognitosrp.CognitoSRP
	cognito *cognitoidentityprovider.Client

	authTokenCache *types.AuthenticationResultType

	client *http.Client
}

func New(cfg *Config) (*Client, error) {
	csrp, err := cognitosrp.NewCognitoSRP(cfg.APIKey, cfg.APISecret, cfg.UserPoolID, cfg.ClientPoolID, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:  cfg,
		csrp: csrp,
		cognito: cognitoidentityprovider.New(cognitoidentityprovider.Options{
			Region: "us-west-2",
		}),
		client: http.DefaultClient,
	}, nil
}

func (c *Client) Init(ctx context.Context) error {
	return c.authenticate(ctx)
}

func (c *Client) authenticate(ctx context.Context) error {
	auth, err := c.cognito.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_SRP_AUTH",
		ClientId:       aws.String(c.csrp.GetClientId()),
		AuthParameters: c.csrp.GetAuthParams(),
	})
	if err != nil {
		return err
	}

	if auth.ChallengeName == types.ChallengeNameTypePasswordVerifier {
		challengeResponses, _ := c.csrp.PasswordVerifierChallenge(auth.ChallengeParameters, time.Now())

		resp, err := c.cognito.RespondToAuthChallenge(
			ctx,
			&cognitoidentityprovider.RespondToAuthChallengeInput{
				ChallengeName:      types.ChallengeNameTypePasswordVerifier,
				ChallengeResponses: challengeResponses,
				ClientId:           aws.String(c.csrp.GetClientId()),
			})
		if err != nil {
			return err
		}

		c.authTokenCache = resp.AuthenticationResult

		return nil
	} else {
		return fmt.Errorf("invalid authentication challenge %v", auth.ChallengeName)
	}
}

func (c *Client) createReq(ctx context.Context, method, route string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%v/%v", c.cfg.APIURL, route), body)
	if err != nil {
		return nil, err
	}

	if c.authTokenCache == nil || c.authTokenCache.AccessToken == nil {
		return nil, fmt.Errorf("client not authenticated, call Client.Init()")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", *c.authTokenCache.AccessToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", c.cfg.APIKey)

	return req, nil
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= 300 {
		msg := new(ErrorMsg)
		err = json.NewDecoder(resp.Body).Decode(msg)
		if err != nil {
			raw := new(json.RawMessage)
			_ = json.NewDecoder(resp.Body).Decode(&raw)
			return resp, fmt.Errorf("HTTP Error [status code=%v] [body=%v]", resp.StatusCode, string(*raw))
		}
		return resp, fmt.Errorf("HTTP Error [status code=%v] [message=%v]", resp.StatusCode, msg.Message)
	}

	return resp, err
}
