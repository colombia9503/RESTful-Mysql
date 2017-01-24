package common

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
)

type AppClaims struct {
	Usuario string `json:"usuario"`
	Nombre  string `json:"nombre"`
	Rol     string `json:"rol"`
	jwt.StandardClaims
}

const (
	privKeyPath = "keys/tm.rsa"
	pubKeyPath  = "keys/tm.rsa.pub"
)

var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//Read key files
func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

func GenerateJWT(user, name, role string) (string, error) {
	//Create the claims
	claims := AppClaims{
		user,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}
	}

	if token.Valid {
		// Set user name to HTTP context
		context.Set(r, "user", token.Claims.(*AppClaims).Usuario)
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}

//Refresh token, taking the claims and the expiration time
func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	/**rtoken, err := TokenFromAuthHeader(r)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"Can't decode the token",
			http.StatusBadGateway,
		)
	}*/
}

//LogOut, expiring the token in request
func LogOut(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
