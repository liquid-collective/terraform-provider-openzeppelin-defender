package defenderclienthttp

import (
	"context"
	"testing"

	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
	"github.com/stretchr/testify/require"
)

func newClient() (*Client, error) { //nolint
	client, err := New((&Config{
		APIKey:    "<placeholder>",
		APISecret: "<placeholder>",
	}).SetDefault())

	if err != nil {
		return nil, err
	}

	err = client.Init(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func testCreateSendUSDCProposal(t testing.TB) { //nolint
	client, err := newClient()
	require.NoError(t, err)

	reqMsg := &defenderclient.CreateProposalReqMsg{
		Via:     "0xE3208Aa9d1186c1D1C8A5b76E794b2B68E6cb3a5",
		ViaType: "Gnosis Safe",
		Contract: defenderclient.ContractMsg{
			Network: "mainnet",
			Address: "0xE3208Aa9d1186c1D1C8A5b76E794b2B68E6cb3a5",
		},
		FunctionInterface: defenderclient.FunctionInterface{
			Name:   "",
			Inputs: []defenderclient.Inputs{},
		},
		FunctionInputs: []interface{}{},
		Type:           "send-funds",
		Metadata: &defenderclient.ProposalMetadata{
			SendTo:    "0x5B453a19845e7492ee3A0df4Ef085D4C75e5752b",
			SendValue: "325000000000",
			SendCurrency: &defenderclient.Currency{
				Address:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Decimals: 6,
				Name:     "USD Coin",
				Network:  "goerli",
				Symbol:   "USDC",
				Type:     "ERC20",
			},
		},
	}

	_, err = client.CreateProposal(context.TODO(), reqMsg)
	require.NoError(t, err)
}
