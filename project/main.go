package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "12345678901234567890"
	b, _ := regexp.MatchString("^[0-9]+$", str)

	fmt.Println(b)
}
