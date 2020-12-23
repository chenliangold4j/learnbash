package main

import (
	"fmt"
	"os"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {

	fmt.Println(os.Args[0])
	if os.Args[0] == "/proc/self/exe" {
		fmt.Println("test")
	} else {
		fmt.Println("kkk")
	}

}
