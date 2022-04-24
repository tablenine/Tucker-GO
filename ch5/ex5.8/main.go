package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(err)
		s, _ := stdin.ReadString('\n')
		fmt.Println(s)
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
