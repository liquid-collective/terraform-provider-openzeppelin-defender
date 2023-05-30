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

	ContractID        string            `json:"contractId"`
	Contract          ContractMsg       `json:"contract,omitempty"`
	FunctionInterface FunctionInterface `json:"targetFunction"`
	FunctionInputs    []interface{}     `json:"functionInputs"`

	IsActive   bool `json:"isActive"`
	IsArchived bool `json:"isArchived"`

	Metadata *ProposalMetadata `json:"metadata,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

type ProposalMetadata struct {
	NewImplementationAddress string    `json:"newImplementationAddress,omitempty"`
	NewImplementationABI     string    `json:"newImplementationAbi,omitempty"`
	ProxyAdminAddress        string    `json:"proxyAdminAddress,omitempty"`
	Action                   string    `json:"action,omitempty"`
	OperationType            string    `json:"operationType,omitempty"`
	Account                  string    `json:"account,omitempty"`
	Role                     string    `json:"role,omitempty"`
	SendTo                   string    `json:"sendTo,omitempty"`
	SendValue                string    `json:"sendValue,omitempty"`
	SendCurrency             *Currency `json:"sendCurrency,omitempty"`
}

type Currency struct {
	Name     string `json:"name,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Address  string `json:"address,omitempty"`
	Network  string `json:"network,omitempty"`
	Decimals int    `json:"decimals,omitempty"`
	Type     string `json:"type,omitempty"`
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

	Contract          ContractMsg        `json:"contract"`
	TargetFunction    *FunctionInterface `json:"targetFunction,omitempty"`
	FunctionInterface *FunctionInterface `json:"functionInterface,omitempty"`
	FunctionInputs    []interface{}      `json:"functionInputs"`

	Metadata *ProposalMetadata `json:"metadata,omitempty"`
}

type Client interface {
	CreateProposal(ctx context.Context, msg *CreateProposalReqMsg) (*ProposalRespMsg, error)
	ListProposals(ctx context.Context, includeArchived bool) ([]*ProposalRespMsg, error)
	GetProposalByList(ctx context.Context, contractID, proposalID string) (*ProposalRespMsg, error)
	ArchiveProposal(ctx context.Context, contractID, proposalID string) error
}
