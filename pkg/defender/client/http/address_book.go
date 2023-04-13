package defenderclienthttp

import (
	"context"
)

func (c *Client) AddressBook(_ context.Context) (map[string]string, error) {
	return nil, ErrNotImplemented
}
