package defenderclienthttp

import (
	"context"
	"encoding/json"
	"net/http"

	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
)

func (c *Client) ListContracts(ctx context.Context) ([]*defenderclient.ContractMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, "contracts", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var msgs []*defenderclient.ContractMsg
	err = json.NewDecoder(resp.Body).Decode(&msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
