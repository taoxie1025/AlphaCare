package lib

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	encrptionCost = bcrypt.DefaultCost
)

func SaltString(s string) (string, error) {
	bArr := []byte(s)
	hash, err := bcrypt.GenerateFromPassword(bArr, encrptionCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashedStrings(hashed, plain string) bool {
	hashedArr, plainArr := []byte(hashed), []byte(plain)
	err := bcrypt.CompareHashAndPassword(hashedArr, plainArr)

	return err == nil // if no error then hashes are the same.
}
