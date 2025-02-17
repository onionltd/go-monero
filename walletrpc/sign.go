package walletrpc

type SignRequest struct {
	// Anything you need to sign.
	Data string `json:"data"`
}

type SignResponse struct {
	// Signature generated against the "data" and the account public address.
	Signature string `json:"signature"`
}

// Sign a string.
func (c *Client) Sign(req *SignRequest) (*SignResponse, error) {
	resp := &SignResponse{}
	err := c.Do("sign", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
