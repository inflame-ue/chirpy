package auth

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "helloWorld12345"
	hash, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}

	if !match {
		t.Errorf("expected %v, got %v", true, false)
	}

}

func TestValidateJWT(t *testing.T) {
	tokenSecret := "testing-secret-token"
	expectedUserID := uuid.New()
	token, err := MakeJWT(expectedUserID, tokenSecret, time.Duration(10*time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	userID, err := ValidateJWT(token, tokenSecret)
	if err != nil {
		log.Fatal(err)
	}

	if userID != expectedUserID {
		t.Errorf("expected: %v, got %v", expectedUserID, userID)
	}
}

func TestValidateJWTExpired(t *testing.T) {
	tokenSecret := "testing-secret-token"
	expectedUserID := uuid.New()
	token, err := MakeJWT(expectedUserID, tokenSecret, time.Duration(1*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second) // sleep for a second to make sure the token expires
	_, err = ValidateJWT(token, tokenSecret)
	if err == nil {
		t.Errorf("expected the token to expire")
	}
}
