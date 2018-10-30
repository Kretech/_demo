package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Kretech/framework/config"
	. "github.com/Kretech/framework/foundation"
	"github.com/Kretech/framework/foundation/console"
)

func init() {
	app := App()

	app.RegisterService(config.NewTomlFileServiceRegister("conf.toml"))

	app.Singleton(`console`, console.NewKernel())
}

func main() {

	console := App().Make(`console`).(*console.Kernel)

	console.RegisterCommand(&SayHiCmd{})
	console.RegisterCommand(&HiWebCmd{})

	out := bytes.NewBuffer([]byte{})
	console.Handle(os.Args, out)

	fmt.Print(out)
}
