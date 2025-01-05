package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/mattn/go-sqlite3"
)

func generateSalt() string {
	b := make([]byte, 32)
	s1 := rand.NewSource(time.Now().UnixNano()) // 根據取得的時間（ns），當成一個種子碼
	r1 := rand.New(s1)                          // 使用種子碼生成隨機數
	r1.Read(b)
	return fmt.Sprintf("%x", b)
}
func hashPassword(password string) (string, string, error) {
	salt := generateSalt()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashedPassword), salt, nil // return hashed password, salt value, error message
}
func verifyPassword(password string, storedHash string, salt string) bool {
	newHash, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(newHash) == storedHash
}
