package main

import (
	"io"

	consoleContract "github.com/Kretech/framework/foundation/console"
)

type SayHiCmd struct {
	consoleContract.BaseCommand
}

func (*SayHiCmd) Name() string {
	return `sayhi`
}

func (*SayHiCmd) Note() string {
	return `say hello`
}

func (*SayHiCmd) Handle(args []string, w io.Writer) {
	w.Write([]byte("hi, this is a demo program for Command\n"))
}
