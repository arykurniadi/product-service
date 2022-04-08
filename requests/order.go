package requests

type (
	OrderCreate struct {
		TransactionNumber string        `json:"transaction_number"`
		Media             string        `json:"media"`
		IsMember          bool          `json:"is_member"`
		Customer          CustomerOrder `json:"customer"`
		OrderItems        []OrderItem   `json:"items"`
	}

	CustomerOrder struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	OrderItem struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	}
)
