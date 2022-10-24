package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
)

// Use this resource to create a proposal
func resourceProposal() *schema.Resource {
	return &schema.Resource{
		CreateContext: proposalCreate,
		ReadContext:   proposalRead,
		UpdateContext: proposalUpdate,
		DeleteContext: proposalDelete,
		Description: `Resource is used to managed Defender Admin action proposals. 
Any actions created this way will have no approvals initially.`,
		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Title of the proposal",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the proposal",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the proposal",
			},
			"via": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Address of the multisig via which the request is sent",
			},
			"via_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the multisig",
			},
			"contract_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the contract on Defender",
			},
			"contract_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the contract as set on Defender when creating the contract",
			},
			"contract_network": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network on which to execute the proposal on",
			},
			"contract_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Address of the contract to execute the proposal on",
			},
			"contract_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"function_interface_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ABI function name to call",
			},
			"function_interface_inputs": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "ABI inputs of the function to call",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"function_inputs": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Inputs of the function to call",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"function_inputs_json": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"function_inputs"},
				Description:   "Inputs of the function to call (in JSON format)",
			},
			"is_active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the proposal is active",
			},
			"is_archived": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the proposal has been archived",
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Metadata",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date when the proposal was created",
			},
		},
	}
}

func proposalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := (meta).(defenderclient.Client)

	proposal, err := client.GetProposalByList(
		ctx,
		d.Get("contract_id").(string),
		d.Id(),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("title", proposal.Title)
	d.Set("description", proposal.Description)
	d.Set("type", proposal.Type)
	d.Set("via", proposal.Via)
	d.Set("via_type", proposal.ViaType)
	d.Set("contract_id", proposal.Contract.ContractID)
	d.Set("contract_name", proposal.Contract.Name)
	d.Set("contract_network", proposal.Contract.Network)
	d.Set("contract_address", proposal.Contract.Address)
	d.Set("contract_type", proposal.Contract.Type)
	d.Set("function_interface_name", proposal.FunctionInterface.Name)
	d.Set("function_interface_inputs", flattenFunctionInterfaceInputs(proposal.FunctionInterface.Inputs))
	err = d.Set("function_inputs", proposal.FunctionInputs)
	if err != nil {
		b, err := json.Marshal(proposal.FunctionInputs)
		if err != nil {
			return diag.FromErr(fmt.Errorf("invalid function inputs %v", err))
		}
		d.Set("function_inputs_json", string(b))
	}
	d.Set("is_active", proposal.IsActive)
	d.Set("is_archived", proposal.IsArchived)
	d.Set("created_at", proposal.CreatedAt.String())
	d.Set("metadata", proposal.Metadata)

	return nil
}

func proposalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := (meta).(defenderclient.Client)

	createProposalReq := &defenderclient.CreateProposalReqMsg{
		Title:       d.Get("title").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Via:         d.Get("via").(string),
		ViaType:     d.Get("via_type").(string),
		Contract: defenderclient.ContractMsg{
			Network: d.Get("contract_network").(string),
			Address: d.Get("contract_address").(string),
		},
		FunctionInterface: defenderclient.FunctionInterface{
			Name:   d.Get("function_interface_name").(string),
			Inputs: []defenderclient.Inputs{},
		},
		FunctionInputs: d.Get("function_inputs").([]interface{}),
		Metadata:       make(map[string]string),
	}

	inputs := d.Get("function_interface_inputs").([]interface{})
	for _, input := range inputs {
		i := input.(map[string]interface{})
		createProposalReq.FunctionInterface.Inputs = append(
			createProposalReq.FunctionInterface.Inputs,
			defenderclient.Inputs{
				Name: i["name"].(string),
				Type: i["type"].(string),
			},
		)
	}

	functionInputsJSON := d.Get("function_inputs_json").(string)
	if functionInputsJSON != "" {
		var functionInputs []json.RawMessage
		err := json.Unmarshal([]byte(functionInputsJSON), &functionInputs)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, functionInput := range functionInputs {
			createProposalReq.FunctionInputs = append(createProposalReq.FunctionInputs, functionInput)
		}
	}

	metadata := d.Get("metadata").(map[string]interface{})
	for k, v := range metadata {
		createProposalReq.Metadata[k] = v.(string)
	}

	proposal, err := client.CreateProposal(ctx, createProposalReq)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(proposal.ProposalID)
	d.Set("contract_id", proposal.Contract.ContractID)

	return proposalRead(ctx, d, meta)
}

func proposalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func proposalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := (meta).(defenderclient.Client).ArchiveProposal(
		ctx,
		d.Get("contract_id").(string),
		d.Id(),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func flattenFunctionInterfaceInputs(funcInputs []defenderclient.Inputs) []map[string]string {
	inputs := []map[string]string{}
	for _, input := range funcInputs {
		inputs = append(inputs, map[string]string{
			"name": input.Name,
			"type": input.Type,
		})
	}

	return inputs
}
