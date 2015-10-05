package main

import (
	"fmt"

	"github.com/oneplus1000/trymssql/core"
)

func main() {
	err := core.SelectFromMsSql()
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
}
