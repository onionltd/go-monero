package walletrpc

type SetAccountTagDescriptionRequest struct {
	// Set a description for this tag.
	Tag string `json:"tag"`

	// Description for the tag.
	Description string `json:"description"`
}

// Set description for an account tag.
func (c *Client) SetAccountTagDescription(req *SetAccountTagDescriptionRequest) error {
	err := c.Do("set_account_tag_description", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
