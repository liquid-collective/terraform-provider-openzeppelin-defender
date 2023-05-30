package defenderclienthttp

import (
	"context"
	"os"
	"testing"

	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
	"github.com/stretchr/testify/require"
)

func newClient() (*Client, error) { //nolint
	client, err := New((&Config{
		APIKey:    os.Getenv("DEFENDER_API_KEY"),
		APISecret: os.Getenv("DEFENDER_API_SECRET"),
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

func testCreateAddOperatorProposal(t *testing.T) { //nolint
	client, err := newClient()
	require.NoError(t, err)

	reqMsg := &defenderclient.CreateProposalReqMsg{
		Title:   "Test (to be archived)",
		Via:     "0xE3208Aa9d1186c1D1C8A5b76E794b2B68E6cb3a5",
		ViaType: "Gnosis Safe",
		Contract: defenderclient.ContractMsg{
			Network: "mainnet",
			Address: "0xAF5ae0768291C4ADC57Cb7e1482F336b4Faa4291",
		},
		TargetFunction: &defenderclient.FunctionInterface{
			Name:   "",
			Inputs: []defenderclient.Inputs{},
		},
		FunctionInterface: &defenderclient.FunctionInterface{
			Name: "addOperator",
			Inputs: []defenderclient.Inputs{
				{
					Name: "_name",
					Type: "string",
				},
				{
					Name: "_operator",
					Type: "address",
				},
			},
		},
		FunctionInputs: []interface{}{"test", "0x0000000000000000000000000000000000000000"},
		Type:           "custom",
	}

	_, err = client.CreateProposal(context.TODO(), reqMsg)
	require.NoError(t, err)
}

func testCreateSendUSDCProposal(t *testing.T) { //nolint
	client, err := newClient()
	require.NoError(t, err)

	reqMsg := &defenderclient.CreateProposalReqMsg{
		Title:   "Test (to be archived)",
		Via:     "0xE3208Aa9d1186c1D1C8A5b76E794b2B68E6cb3a5",
		ViaType: "Gnosis Safe",
		Contract: defenderclient.ContractMsg{
			Network: "mainnet",
			Address: "0xE3208Aa9d1186c1D1C8A5b76E794b2B68E6cb3a5",
		},
		TargetFunction: &defenderclient.FunctionInterface{
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
