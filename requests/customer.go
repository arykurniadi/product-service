package requests

type (
	CustomerCreate struct {
		Name    string `form:"name"`
		Email   string `form:"email"`
		Phone   string `form:"phone"`
		Address string `form:"address"`
	}
)
