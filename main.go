package main

import (
	"fmt"

	v1 "github.com/lakhansamani/go-test-lib/pkg/v1"
	v2 "github.com/lakhansamani/go-test-lib/pkg/v2"
)

func main() {
	v1Sum := v1.Sum(5, 10)
	fmt.Println("v1 sum", v1Sum)
	v2Sum := v2.Sum(5, 10, 15)
	fmt.Println("v2 sum", v2Sum)
}
