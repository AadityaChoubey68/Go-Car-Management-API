package login

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AadityaChoubey68/Go-Car-Management-API/models"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creadentials models.Creadentials
	if err := json.NewDecoder(r.Body).Decode(&creadentials); err != nil {
		http.Error(w, "Error Decoding Credentials", http.StatusBadRequest)
		return
	}

	valid := (creadentials.Username == "admin" && creadentials.Password == "admin12345")
	if !valid {
		http.Error(w, "Incorrect Username or Password", http.StatusUnauthorized)
		return
	}

	tokenString, err := Generatetoken(creadentials.Username)
	if err != nil {
		http.Error(w, "Failed to Generate token", http.StatusInternalServerError)
		log.Println("Error generating token")
		return
	}

	response := map[string]string{"token": tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Generatetoken(userName string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("some_value"))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}
