package main

import (
	"fmt"
	
	"github.com/flan0910/SimpleProxy/modules"
)

func main() {
	modules.LoadConfig()
	fmt.Println(modules.GetConfig())
}