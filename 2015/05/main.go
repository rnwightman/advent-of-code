package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if IsNice(line) {
			c += 1
		}
	}

	fmt.Printf("# of Nice: %d\n", c)
}
