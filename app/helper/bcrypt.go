package helper

import (
	"depobangunan/app/environment"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	environment.ExportEnv()

	keypass := os.Getenv("KEYPASS")

	convertKeypass, err := strconv.Atoi(keypass)
	if err != nil {
		return "failed convert string to integer", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), convertKeypass)
	if err != nil {
		return "error hashpassword", err
	}

	return string(hashedPassword), nil
}

func DecryptPassword(hashedPassword string, passwordInput string) (bool) {
	

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordInput))
	if err != nil {
		fmt.Println("Passwords do not match.")
		return false
	}

	return true
}
