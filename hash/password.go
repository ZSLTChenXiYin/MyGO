package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHasher struct{}

func NewPasswordHasher() *PasswordHasher { return &PasswordHasher{} }

func (ph *PasswordHasher) Hash(password string) (string, error) {
	password_hash_bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(password_hash_bytes), err
}

func (ph *PasswordHasher) Compare(password string, password_hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
}
