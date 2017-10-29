package main

import (
	"fmt"

	"bytes"

	. "github.com/Kretech/framework/foundation"
	"github.com/Kretech/framework/foundation/console"
)

func init() {
	app := App()
	app.Singleton(`console`, console.NewKernel())
}

func main() {
	console := App().Make(`console`).(*console.Kernel)
	console.RegisterCommand(&SayHiCmd{})

	out := bytes.NewBuffer([]byte{})
	console.Handle(out)

	fmt.Print(out)
}
