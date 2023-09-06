package functions

import "fmt"

func print(printing string) {
	fmt.Print(printing)
}

func println(printing string) {
	fmt.Println(printing)
}

func scanner(scanning string) {
	fmt.Scanf("%s", &scanning)
	println(scanning)
}
