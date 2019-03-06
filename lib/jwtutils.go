package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	conf "../config"
	jwt "github.com/dgrijalva/jwt-go"
)

//ConvertToTicketNumber to convert data to ticket number
func GetToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var mySigningKey = []byte(conf.Param.TokenAuth)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}
func TokenAuth(mdn string, key string, time string) bool {
	var secrectKey = conf.Param.TokenAuth

	h := sha256.New()
	h.Write([]byte(secrectKey + mdn + time))
	ecriptkey := hex.EncodeToString(h.Sum(nil))
	if key == ecriptkey {
		return true
	}
	return false
}
