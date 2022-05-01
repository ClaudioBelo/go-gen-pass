package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
)

var lowerCharSet = "abcdedfghijklmnopqrst"
var upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numberSet = "01234567890123456789"
var normalCharSet = lowerCharSet + upperCharSet + numberSet
var specialCharSet = "!@#$%&*!@#$%&*"
var allCharSet = normalCharSet + specialCharSet

func main() {
	length, err := strconv.Atoi(os.Args[1])
	charSet := allCharSet
	if (len(os.Args) > 2) &&
		(os.Args[2] == "-ns" || os.Args[2] == "--no-special") ||
		(os.Args[1] == "-ns" || os.Args[1] == "--no-special") {
		charSet = normalCharSet
	}
	if err != nil {
		length = 20
	}
	password := genPassword(length, charSet)
	if err := clipboard.WriteAll(password); err != nil {
		panic(err)
	}
	fmt.Println(password)
	fmt.Println("copied to clipboard!!")
}

func genPassword(length int, charSet string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var password []byte
	for i := 0; i < length; i++ {
		password = append(password, charSet[rand.Intn(len(charSet))])
	}
	return string(password)
}
