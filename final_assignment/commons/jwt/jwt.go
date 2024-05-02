package jwt

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
	"hacktiv8-course/final_assignment/internal/entities"
	"log"
	"os"
	"time"
)

type JwtPkg struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

type JwtTokenClaim struct {
	jwt.StandardClaims
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func RegisterJwt(privateKeyFile, publicKeyFile, secretKey string) *JwtPkg {
	var (
		privateKeyBytes []byte
		publicKeyBytes  []byte
		err             error
	)

	privateKeyBytes, err = os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	// Parse the private key.
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(privateKeyBytes, secretKey)
	if err != nil {
		log.Fatalf("Failed to parse private key file: %v", err)
	}

	publicKeyBytes, err = os.ReadFile(publicKeyFile)
	if err != nil {
		log.Fatalf("Failed to read public key file: %v", err)
	}

	// Parse the public key.
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	return &JwtPkg{publicKey: publicKey, privateKey: privateKey}
}

func (j JwtPkg) GenerateToken(user entities.User, duration int64) (tokenString string, err error) {

	jwtTokenClaims := JwtTokenClaim{
		StandardClaims: jwt.StandardClaims{
			Id:        ksuid.New().String(),
			Issuer:    "indrahactiv.com",
			IssuedAt:  time.Now().Unix(),
			Subject:   "user",
			Audience:  "indrahactivapps",
			ExpiresAt: time.Now().Add(time.Second * time.Duration(duration)).Unix(),
		},
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// Create a new token with the claims and sign it using the private key.
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtTokenClaims)
	tokenString, err = token.SignedString(j.privateKey)
	if err != nil {
		fmt.Printf("Failed to sign token: %v\n", err)
		return
	}
	return
}

func (j JwtPkg) ValidateToken(tokenString string) (user entities.User, err error) {
	var (
		token          *jwt.Token
		jwtTokenClaims JwtTokenClaim
	)

	// Parse and validate the JWT token using the public key.
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is RSA.
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err = fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.publicKey, nil
	})

	if err != nil {
		log.Printf("Token validation failed: %v", err)
		err = fmt.Errorf("Token validation failed: %v", err)
		return
	}

	switch token.Valid {
	case true:
		// Token is valid. You can access its claims.
		claims := token.Claims.(jwt.MapClaims)
		jwtTokenClaims = j.getDataClaims(claims)

		//message response
		user.ID = jwtTokenClaims.Id
		user.Username = jwtTokenClaims.Username

	default:
		err = fmt.Errorf("token is invalid")
	}

	return

}

func (j JwtPkg) getDataClaims(claims jwt.MapClaims) (jwtTokenClaims JwtTokenClaim) {

	jwtTokenClaims.Id = uint(claims["id"].(float64))
	jwtTokenClaims.Username = claims["username"].(string)

	return
}
