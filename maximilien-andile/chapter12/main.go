// package-init/rules-illustration/main.go
package main

import (
	"fmt"

	"maximilien-andile.com/packageInit/rules/invoice"
)

func init() {
	fmt.Println("main")
}
func main() {
	fmt.Println("--program start--")
	invoice.Print()
}
