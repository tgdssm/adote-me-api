package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

func CheckPasswd(passwdWithHash, passwdWithoutHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwdWithHash), []byte(passwdWithoutHash))
}
