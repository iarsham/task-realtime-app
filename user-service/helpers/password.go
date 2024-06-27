package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPass(plainPass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
}

func ValidatePass(hashedPass string, plainPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
}
