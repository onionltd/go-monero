package walletrpc

type GetPaymentsRequest struct {
	// Payment ID used to find the payments (16 characters hex).
	PaymentId string `json:"payment_id"`
}

type GetPaymentsResponse struct {
	// List of payments
	Payments []Payment `json:"payments"`
}

// Get a list of incoming payments using a given payment id.
func (c *Client) GetPayments(req *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	resp := &GetPaymentsResponse{}
	err := c.Do("get_payments", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
