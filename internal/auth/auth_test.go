package auth

import (
	"log"
	"testing"
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
