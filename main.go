package main

import (
	"fmt"

	_ "github.com/go-xorm/xorm"
	_ "k8s.io/apimachinery"
)

func main() {
	fmt.Println("Hello world!")
}
