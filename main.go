// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str, err := getInputString()

	if err != nil {
		fail(err)
	}

	num := countWords(str)
	fmt.Println(num)
}

func getInputString() (string, error) {
	if len(os.Args) == 2 {
		return os.Args[1], nil
	} else if len(os.Args) == 1 {
		return "", nil
	} else {
		return "", fmt.Errorf("wrong arguments")
	}
}

func countWords(str string) int {
	s := strings.Fields(str)
	return len(s)
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
