package main

import (
	"fmt"
	"github.com/crafty-ezhik/observa-core/internal/app"
)

func main() {
	dep := app.Bootstrap()
	fmt.Println(dep)

}
