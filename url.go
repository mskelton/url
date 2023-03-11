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
		`(https?|ftp):\/\/` + // scheme
			`(\.?[a-zA-Z0-9_\-])+` + // subdomain
			`(\.[a-zA-Z]+)?` + // domain
			`(:[0-9]+)?` + // port
			`(\.?[a-zA-Z0-9%_=~/+\-@]+)*` + // path
			`(\?([.,]?[a-zA-Z0-9&%_=~+\-@])*)?` + // query
			`(#[a-zA-Z0-9&%_=~+\-@]*)?`, // fragment
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
