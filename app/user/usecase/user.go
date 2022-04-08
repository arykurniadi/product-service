package usecase

import (
	UserInterface "dbo.id/product-service/app/user"
	"dbo.id/product-service/libraries"
	"dbo.id/product-service/models"
	"dbo.id/product-service/requests"
	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	UserRepository UserInterface.IUserRepository
}

var jwtService libraries.JWTService = libraries.JWTAuthService()

func NewUserUsecase(u UserInterface.IUserRepository) UserInterface.IUserUsecase {
	return &UserUsecase{
		UserRepository: u,
	}
}

func (a *UserUsecase) GetListUser(c *gin.Context, page int, perPage int) (users []models.User, pagination *models.Pagination, err error) {
	users, pagination, err = a.UserRepository.GetListUser(page, perPage)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (a *UserUsecase) Create(c *gin.Context, req requests.UserCreate) (user *models.User, err error) {
	user = &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
		Token:    jwtService.GenerateToken(req.Email, true),
	}

	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYXJ5a3VybmlhZGlAZ21haWwuY29tIiwidXNlciI6dHJ1ZSwiZXhwIjo1MjU2MDAwMDAwLCJpYXQiOjE2NDkzOTI5NDksImlzcyI6IkpXVENyZWF0ZSJ9.nMqnbc3t1VW5BwSYof6nX5ryckKbZMKnHtcJzBEwB5w"
	// validate, err := jwtService.ValidateToken(token)
	// if validate.Valid {
	// 	claims := validate.Claims.(jwt.MapClaims)
	// 	fmt.Println(claims)
	// }

	res, err := a.UserRepository.Create(*user)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *UserUsecase) GetUserById(c *gin.Context, id int) (user models.User, err error) {
	user, err = a.UserRepository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (a *UserUsecase) Update(c *gin.Context, id int, req requests.UserUpdate) (user *models.User, err error) {
	user = &models.User{
		Id:       id,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
	}

	res, err := a.UserRepository.Update(id, *user)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *UserUsecase) Delete(c *gin.Context, id int) (user models.User, err error) {
	user, err = a.UserRepository.GetUserById(id)
	if err != nil {
		return user, err
	}

	err = a.UserRepository.Delete(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
