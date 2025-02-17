package walletrpc

type LabelAccountRequest struct {
	// Apply label to account at this index.
	AccountIndex uint64 `json:"account_index"`

	// Label for the account.
	Label string `json:"label"`
}

// Label an account.
func (c *Client) LabelAccount(req *LabelAccountRequest) error {
	err := c.Do("label_account", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
