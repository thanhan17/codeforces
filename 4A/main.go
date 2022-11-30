package main

import "fmt"

func main() {
	var w int
	fmt.Scanf("%d", &w)
	if w%2 == 0 && w > 3 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
