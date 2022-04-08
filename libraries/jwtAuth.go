package libraries

import (
	"fmt"
	"time"

	"dbo.id/product-service/config"
	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

var appConfig = config.Config

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		// issure:    "Bikash",
		issure: "JWTCreate",
	}
}

func getSecretKey() string {
	secret := appConfig.JWT.Secret
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			ExpiresAt: int64(appConfig.JWT.Ttl),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}

// func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
// 	tokenString := strings.Split(encodedToken, " ")
// 	fmt.Println("validateToken", tokenString[1])

// 	return nil, nil
// }
