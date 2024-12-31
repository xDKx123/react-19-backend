package middleware

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, salt string) (string, error) {
	combinedPassword := password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(combinedPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, salt string, userHashedPassword string) bool {
	combinedPassword := password + salt
	return bcrypt.CompareHashAndPassword([]byte(userHashedPassword), []byte(combinedPassword)) == nil
}

func GenerateSalt() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte{}, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	salt := string(hash[:16])

	return salt, nil
}
