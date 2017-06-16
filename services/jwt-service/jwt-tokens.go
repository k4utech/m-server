package jts

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/k4utech/m-server/models"
)

const signingkey = "IKnewYouWillSeeMeeAroundHello###WelcomeToTheClub:)"

// ExampleCreatingToken : Test Method to test working
func ExampleCreatingToken() {

	token := jwt.New(jwt.SigningMethodHS256) //Creating JWT Token

	//Adding claims
	claims := make(jwt.MapClaims)
	claims["user"] = "true"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims

	//Sighning the token
	signedToken, ok := token.SignedString([]byte(signingkey))

	if ok != nil {
		fmt.Println([]byte(ok.Error()))
	} else {
		fmt.Println([]byte(signedToken))
	}
}

// ExampleVerifyingToken : Test Method to to verify a JWT Tokens
func ExampleVerifyingToken() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc4OTc1NjYsIm5hbWUiOiJ0cnVlIiwidXNlciI6InRydWUifQ.OE3v4g-_Grx18htmaQAOt2y7udW_RUuvv1nELK4Z-Sk"
	ValidateToken(token)
}

// CreateToken : Generate a token for given user
func CreateToken(user model.User) (string, error) {
	//Creating token
	token := jwt.New(jwt.SigningMethodHS256)

	//Adding claims
	claims := make(jwt.MapClaims)
	claims["name"] = user.Name
	claims["mobile"] = user.Mobile
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims

	//Signing the token
	signedToken, ok := token.SignedString([]byte(signingkey))

	return signedToken, ok
}

func parseToken(signedtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedtoken, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	return token, err
}

//ValidateTokenAndGetClaims : Validate the token that it belongs to a perticular user and not expired yet
func ValidateTokenAndGetClaims(signedtoken string) (jwt.MapClaims, error) {

	valid, unsignedToken, err := ValidateToken(signedtoken)

	if valid == false {
		return nil, err
	}

	mapClamins := unsignedToken.Claims.(jwt.MapClaims) //Retrive claims map

	return mapClamins, nil
}

//ValidateToken : Validate the token that it belongs to a perticular user and not expired yet
func ValidateToken(signedtoken string) (bool, *jwt.Token, error) {

	token, err := parseToken(signedtoken) //Retrive decoded token

	if err != nil || !token.Valid {
		fmt.Println(err.Error())
		return false, token, err
	}

	return true, token, err
}
