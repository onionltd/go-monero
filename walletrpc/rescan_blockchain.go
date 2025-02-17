package walletrpc

// Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself.
func (c *Client) RescanBlockchain() error {
	err := c.Do("rescan_blockchain", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
