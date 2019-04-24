package main

import (
	"fmt"
	"github.com/gogf/gf"
	"github.com/gogf/gf/g"
	_ "my_test_modules/boot"
	_ "my_test_modules/router"
)

func main() {

	fmt.Print("hello", gf.VERSION)
	g.Server().Run()
}
