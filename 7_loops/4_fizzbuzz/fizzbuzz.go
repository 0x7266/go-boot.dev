package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Printf("%d\n", i)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
