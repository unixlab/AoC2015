package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	counter := 0
	input := "iwrupvqb"

	for {
		hasher := md5.New()
		io.WriteString(hasher, fmt.Sprintf("%s%d", input, counter))
		hash := fmt.Sprintf("%x", hasher.Sum(nil))
		if strings.HasPrefix(hash, "00000") {
			fmt.Println(counter)
			fmt.Println(hash)
			break
		}
		counter++
	}
}
