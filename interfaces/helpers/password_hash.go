package helpers

import "golang.org/x/crypto/bcrypt"

func CompareHashAndPassword(password []byte, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
}

func GenerateHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
