package hdlrs

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["name"] = "true"
	claims["user"] = "true"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = claims

	signedToken, ok := token.SignedString([]byte("IKnewYouWillSeeMeeAroundHello###WelcomeToTheClub:)"))

	if ok != nil {
		w.Write([]byte(ok.Error()))
	} else {
		w.Write([]byte(signedToken))
	}
})
