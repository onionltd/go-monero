package walletrpc

type GetAddressBookRequest struct {
	// Indices of the requested address book entries.
	Entries []uint64 `json:"entries"`
}

type GetAddressBookResponse struct {
	// Array of entries.
	Entries []Entry `json:"entries"`
}

// Retrieves entries from the address book.
func (c *Client) GetAddressBook(req *GetAddressBookRequest) (*GetAddressBookResponse, error) {
	resp := &GetAddressBookResponse{}
	err := c.Do("get_address_book", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
