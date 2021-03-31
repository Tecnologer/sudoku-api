package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/tecnologer/go-secrets"
	"github.com/tecnologer/sudoku/clients/sudoku-api/auth"
	"github.com/tecnologer/sudoku/clients/sudoku-api/data"
	"golang.org/x/crypto/bcrypt"
)

var secretkey string

func init() {
	secretkey = secrets.GetKeyString("jwt.key")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection := data.GetDatabase()
	defer data.Closedatabase(connection)

	var user auth.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		internalErrorf(&w, "Error in reading body: %v", err)
		return
	}
	var dbuser auth.User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		preconditionFailedf(&w, "Email already in use: %v", err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	connection.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	connection := data.GetDatabase()
	defer data.Closedatabase(connection)

	var authdetails auth.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		internalErrorf(&w, "Error in reading body: %v", err)
		return
	}

	var authuser auth.User
	connection.Where("email = ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		preconditionFailedf(&w, "Username or Password is incorrect")
		return
	}

	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		preconditionFailedf(&w, "Username or Password is incorrect")
		return
	}

	validToken, err := GenerateJWT(authuser.Email, authuser.Role)
	if err != nil {
		internalErrorf(&w, "Failed to generate token")
		return
	}

	var token auth.Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			unauthorizedErrorf(&w, "No Token Found")
			return
		}

		token, err := getToken(r)
		if err != nil {
			unauthorizedErrorf(&w, "Your Token has been expired")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}

		unauthorizedErrorf(&w, "Not authorized")
	}
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", errors.Wrap(err, "signing string")
	}
	return tokenString, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getToken(r *http.Request) (*jwt.Token, error) {
	var mySigningKey = []byte(secretkey)

	return jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})
}
