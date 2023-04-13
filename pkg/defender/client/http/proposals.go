package defenderclienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client"
)

func (c *Client) CreateProposal(ctx context.Context, msg *defenderclient.CreateProposalReqMsg) (*defenderclient.ProposalRespMsg, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, "proposals", b)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	proposal := new(defenderclient.ProposalRespMsg)
	err = json.NewDecoder(resp.Body).Decode(proposal)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (c *Client) ListProposals(ctx context.Context, includeArchived bool) ([]*defenderclient.ProposalRespMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, "proposals", http.NoBody)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	if includeArchived {
		query.Set("includeArchived", "true")
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var msgs []*defenderclient.ProposalRespMsg
	err = json.NewDecoder(resp.Body).Decode(&msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (c *Client) GetProposal(_ context.Context, _, _ string) (*defenderclient.ProposalRespMsg, error) {
	return nil, ErrNotImplemented
}

func (c *Client) GetProposalByList(ctx context.Context, _, proposalID string) (*defenderclient.ProposalRespMsg, error) {
	proposals, err := c.ListProposals(ctx, true)
	if err != nil {
		return nil, err
	}

	for _, proposal := range proposals {
		if proposal.ProposalID == proposalID {
			return proposal, nil
		}
	}

	return nil, fmt.Errorf("no proposal with proposalId %q", proposalID)
}

type ArchiveReqMsg struct {
	Archived bool `json:"archived"`
}

func (c *Client) ArchiveProposal(ctx context.Context, contractID, proposalID string) error {
	return c.archiveProposal(ctx, contractID, proposalID, &ArchiveReqMsg{true})
}

func (c *Client) archiveProposal(ctx context.Context, contractID, proposalID string, msg *ArchiveReqMsg) error {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(msg)
	if err != nil {
		return err
	}

	req, err := c.createReq(
		ctx,
		http.MethodPut,
		fmt.Sprintf("contracts/%v/proposals/%v/archived", contractID, proposalID),
		b,
	)
	if err != nil {
		return err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	proposal := new(defenderclient.ProposalRespMsg)
	err = json.NewDecoder(resp.Body).Decode(proposal)
	if err != nil {
		return err
	}

	return nil
}
