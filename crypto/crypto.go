package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {}


func (c *Hash) Generate(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func (c *Hash) Compare(hash,password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
