package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numZeros int64
	if len(os.Args) >= 2 {
		numZeros, _ = strconv.ParseInt(os.Args[1], 10, 32)
	}
	if numZeros <= 0 {
		numZeros = 5
	}

	fmt.Printf("Num of Zeros: %d\n", numZeros)

	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')

	fmt.Printf("Secret Key: %s\n", key)
	n := 1
	for {
		hash := hash(key, n)
		if matches(hash, numZeros) {
			break
		}

		n += 1
	}

	fmt.Printf("Lowest Number: %d\n", n)
}

func hash(k string, n int) string {
	var text = fmt.Sprintf("%s%d", k, n)

	bs := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", bs)
}

func matches(h string, numZeros int64) bool {
	for i := 0; i < int(numZeros); i += 1 {
		if h[i] != '0' {
			return false
		}
	}
	return true
}
