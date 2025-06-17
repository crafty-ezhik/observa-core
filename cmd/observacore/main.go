package main

import (
	"fmt"
	"github.com/crafty-ezhik/observa-core/internal/app"
)

func main() {
	dep, appFiber := app.Bootstrap()
	fmt.Println(dep.Db.Password)

	appFiber.Listen(":8080")
}
