package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	str := "0123456789asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP"
	return Random(n, str)
}
func RandomLower(n int) string {
	str := "0123456789asdfghjklzxcvbnmqwertyuiop"
	return Random(n, str)
}

func RandomNumber(n int) string {
	str := "0123456789"
	return Random(n, str)
}

func Random(n int, src string) string {
	bytes := []byte(src)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
