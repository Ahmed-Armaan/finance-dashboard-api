package auth

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Compare(provided []byte, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), provided)
	return err == nil
}
