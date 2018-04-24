package utils

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// FailOnError : avoid some typing when checking for err
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandSeq : Generate random sequence of characters of length n
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Timestamp : convenience function with an expressive name
func Timestamp() time.Time {
	return time.Now().UTC()
}
