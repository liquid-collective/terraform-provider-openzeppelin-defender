package defenderclienthttp

import (
	"context"
)

func (c *Client) AddressBook(ctx context.Context) (map[string]string, error) {
	return nil, ErrNotImplemented

	// TODO: to be implemented when supported by Defender API
	// req, err := c.createReq(ctx, http.MethodGet, "address-book", http.NoBody)
	// if err != nil {
	// 	return nil, err
	// }

	// resp, err := c.client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// defer resp.Body.Close()

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// bodyString := string(bodyBytes)
	// fmt.Printf("%v\n", resp.StatusCode)
	// fmt.Printf("%v\n", bodyString)

	// return nil, nil
}
