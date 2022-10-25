package provider

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
	defenderclientmock "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client/mock"
	"github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/testutils"
)

const testAccResourceCreateProposalConfig = `
resource "defender_proposal" "test1" {
	title = "Test Title"
	description = "Test Description"
	type = "custom"
	via = "0xc0ffee254729296a45a3885639AC7E10F9d54979"
	via_type = "Gnosis Safe"
	contract_network = "goerli"
	contract_address = "0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E"
	function_interface_name = "testMethod"
	function_interface_inputs {
		name = "_testInputAlpha"
		type = "uint256"
	  }
	  function_interface_inputs {
		name = "_testInputBeta"
		type = "address"
	  }
	function_inputs = ["1234","5678"]
}
`

func testProviders(t *testing.T, client defenderclient.Client) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"defender": func() (*schema.Provider, error) {
			t.Setenv("DEFENDER_API_KEY", "test-api-key")
			t.Setenv("DEFENDER_API_SECRET", "test-api-secret")
			provider := Provider()
			provider.ConfigureContextFunc = func(_ context.Context, _ *schema.ResourceData) (interface{}, diag.Diagnostics) {
				return client, nil
			}
			return provider, nil
		},
	}
}

func TestAccResourceProposal(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := defenderclientmock.NewMockClient(ctrl)

	createReqMsg := &defenderclient.CreateProposalReqMsg{
		Title:       "Test Title",
		Description: "Test Description",
		Type:        "custom",
		Via:         "0xc0ffee254729296a45a3885639AC7E10F9d54979",
		ViaType:     "Gnosis Safe",
		Contract: defenderclient.ContractMsg{
			Network: "goerli",
			Address: "0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E",
		},
		FunctionInterface: defenderclient.FunctionInterface{
			Name: "testMethod",
			Inputs: []defenderclient.Inputs{
				{
					Name: "_testInputAlpha",
					Type: "uint256",
				},
				{
					Name: "_testInputBeta",
					Type: "address",
				},
			},
		},
		FunctionInputs: []interface{}{"1234", "5678"},
		Metadata:       make(map[string]string),
	}

	proposalRespMsg := &defenderclient.ProposalRespMsg{
		ProposalID:  "test-proposal-id",
		Title:       "Test Title",
		Description: "Test Description",
		Type:        "custom",
		Via:         "0xc0ffee254729296a45a3885639AC7E10F9d54979",
		ViaType:     "Gnosis Safe",
		Contract: defenderclient.ContractMsg{
			ContractID: "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E",
			Network:    "goerli",
			Address:    "0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E",
			Name:       "Test Contract Name",
			Type:       "Contract",
		},
		FunctionInterface: defenderclient.FunctionInterface{
			Name: "testMethod",
			Inputs: []defenderclient.Inputs{
				{
					Name: "_testInputAlpha",
					Type: "uint256",
				},
				{
					Name: "_testInputBeta",
					Type: "address",
				},
			},
		},
		Metadata:       make(map[string]string),
		CreatedAt:      time.Date(2022, time.October, 24, 23, 13, 0, 0, time.UTC),
		FunctionInputs: []interface{}{"1234", "5678"},
		IsActive:       true,
		IsArchived:     false,
	}

	client.EXPECT().CreateProposal(gomock.Any(), testutils.JSONEq(createReqMsg)).Return(proposalRespMsg, nil)
	client.EXPECT().GetProposalByList(gomock.Any(), "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E", "test-proposal-id").Return(proposalRespMsg, nil).AnyTimes()
	client.EXPECT().ArchiveProposal(gomock.Any(), "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E", "test-proposal-id").Return(nil)
	//	client.EXPECT().GetProposalByList(gomock.Any(), "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E", "test-proposal-id").Return(proposalRespMsg, nil)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(t, client),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCreateProposalConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("defender_proposal.test1", "title", "Test Title"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "description", "Test Description"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "type", "custom"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "via", "0xc0ffee254729296a45a3885639AC7E10F9d54979"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "via_type", "Gnosis Safe"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_address", "0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_id", "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_name", "Test Contract Name"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_type", "Contract"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "created_at", "2022-10-24 23:13:00 +0000 UTC"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.#", "2"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.0", "1234"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.1", "5678"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.#", "2"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.0.name", "_testInputAlpha"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.0.type", "uint256"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.1.name", "_testInputBeta"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.1.type", "address"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_name", "testMethod"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "is_active", "true"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "is_archived", "false"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "metadata.%", "0"),
				),
			},
			{
				Config: testAccResourceCreateProposalConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("defender_proposal.test1", "title", "Test Title"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "description", "Test Description"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "type", "custom"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "via", "0xc0ffee254729296a45a3885639AC7E10F9d54979"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "via_type", "Gnosis Safe"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_address", "0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_id", "goerli-0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_name", "Test Contract Name"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "contract_type", "Contract"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "created_at", "2022-10-24 23:13:00 +0000 UTC"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.#", "2"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.0", "1234"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_inputs.1", "5678"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.#", "2"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.0.name", "_testInputAlpha"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.0.type", "uint256"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.1.name", "_testInputBeta"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_inputs.1.type", "address"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "function_interface_name", "testMethod"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "is_active", "true"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "is_archived", "false"),
					resource.TestCheckResourceAttr("defender_proposal.test1", "metadata.%", "0"),
				),
			},
		},
	})
}
