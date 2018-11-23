package main

import "fmt"
import "math/rand"

var exit bool = false

func testexit() {
	var temp int
	temp = rand.Intn(100)
	exit = (temp == 15)
}

var unused float32

func main() {
	for !exit {
		testexit()
		fmt.Println("still running!")
	}
}
