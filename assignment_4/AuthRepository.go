package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var SecretKey = "U4l1Prekr4sneyshyy$ecretKey"

func Login(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	username := httpParams["username"]
	password := httpParams["password"]

	if !ValidateString(username) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of username may contain SQL injection code")
		return
	}

	if !ValidateString(password) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of password may contain SQL injection code")
		return
	}

	exists, user := UserExists(username)

	if !exists {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("User not found!")
		return
	}

	_, err := CheckPasswordsEquality(user.Password, password)

	if err != nil {
		http.Error(w, http.StatusText(403), http.StatusForbidden)
		log.Println("Incorrect password!")
		return
	}

	token, err := GenerateJWT(user)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Token can not be generated!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
	}

	if !ValidateString(user.Username) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of username may contain SQL injection code")
		return
	}

	if !ValidateString(user.Password) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of password may contain SQL injection code")
		return
	}
	if !ValidateString(user.FullName) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("The content of full name may contain SQL injection code")
		return
	}

	exists, _ := UserExists(user.Username)

	if exists {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		log.Println("User already exists!")
	}

	HashedPassword, err := HashPassword(user.Password)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error occured with password hashing")
		return
	}

	user.Password = HashedPassword

	GetConnection().Create(&user)

}

func UserExists(Username string) (bool, User) {
	var user User
	GetConnection().Find(&user, "username = ?", Username)

	if user == (User{}) {
		return true, user
	} else {
		return false, User{}
	}
}

func HashPassword(RawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RawPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("Failed to hash password: %w", err)
	}

	log.Println("password has successfully been hashed")

	return string(hashedPassword), nil
}

func CheckPasswordsEquality(HashedPassword, RawPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(RawPassword))

	if err != nil {
		return false, fmt.Errorf("Password does not match: %w", err)
	}

	return true, nil
}

func GenerateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId:           user.Id,
		GrantedAuthority: user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		log.Fatal("Error occured with signing token")
		return "", err
	}

	log.Println("Token generated")
	return tokenString, nil
}

func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("User is not authorized")
			return
		}

		tokenString = tokenString[7:] // remove "Bearer " prefix

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("User is not authorized")
			return
		}

		claims, ok := token.Claims.(*Claims)

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("User is not authorized")
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserId)
		ctx = context.WithValue(ctx, "granted_authority", claims.GrantedAuthority)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func CheckAdminRights(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		grantedAuthority := r.Context().Value("granted_authority").(string)

		log.Println("Granted Authority:", grantedAuthority)

		if grantedAuthority != "ADMIN" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
