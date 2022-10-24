package client

import (
	"context"
	"time"
)

type ContractMsg struct {
	ContractID string `json:"contractId,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Network    string `json:"network,omitempty"`
	Address    string `json:"address,omitempty"`
}

type ProposalRespMsg struct {
	ProposalID  string `json:"proposalId"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Type    string `json:"type"`
	Via     string `json:"via"`
	ViaType string `json:"viaType"`

	Contract          ContractMsg       `json:"contract"`
	FunctionInterface FunctionInterface `json:"functionInterface"`
	FunctionInputs    []interface{}     `json:"functionInputs"`

	IsActive   bool `json:"isActive"`
	IsArchived bool `json:"isArchived"`

	Metadata map[string]string `json:"metadata"`

	CreatedAt time.Time `json:"createdAt"`
}

type FunctionInterface struct {
	Name   string   `json:"name"`
	Inputs []Inputs `json:"inputs"`
}

type Inputs struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateProposalReqMsg struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Type    string `json:"type"`
	Via     string `json:"via"`
	ViaType string `json:"viaType"`

	Contract          ContractMsg       `json:"contract"`
	FunctionInterface FunctionInterface `json:"functionInterface"`
	FunctionInputs    []interface{}     `json:"functionInputs"`

	Metadata map[string]string `json:"metadata"`
}

type Client interface {
	CreateProposal(ctx context.Context, msg *CreateProposalReqMsg) (*ProposalRespMsg, error)
	ListProposals(ctx context.Context, includeArchived bool) ([]*ProposalRespMsg, error)
	GetProposalByList(ctx context.Context, contractID, proposalID string) (*ProposalRespMsg, error)
	ArchiveProposal(ctx context.Context, contractID, proposalID string) error
}
