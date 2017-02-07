package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 34
	ages["bob"] = 23
	ages["kim"] = 15

	fmt.Println(ages)
}
