package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func ParseURLs(s string) []string {
	re := regexp.MustCompile(
		`(https?|ftp):\/\/` +
			`(\.?[a-zA-Z0-9_\-])+` +
			`(\.[a-zA-Z]+)?` +
			`(:[0-9]+)?` +
			`(\.?[a-zA-Z0-9%_=~/+-]+)*` +
			`(\?([.,]?[a-zA-Z0-9&%_=~+-])*)?` +
			`(#[a-zA-Z0-9&%_=~+-]*)?`,
	)

	return re.FindAllString(s, -1)
}

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	input := strings.TrimSuffix(string(stdin), "\n")
	for _, url := range ParseURLs(input) {
		fmt.Println(url)
	}
}
