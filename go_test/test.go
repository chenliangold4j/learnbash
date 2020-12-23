package main

import (
	"fmt"
	"go_test/user"
)

func main() {
	fmt.Println("ok")
	var sum = user.Max(1, 2)
	fmt.Println(sum)
}
