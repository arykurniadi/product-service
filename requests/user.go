package requests

type (
	UserCreate struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
		Role     string `form:"role"`
	}

	UserUpdate struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
		Role     string `form:"role"`
	}
)
