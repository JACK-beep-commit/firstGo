package main

import (
	_ "firstGo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"firstGo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
