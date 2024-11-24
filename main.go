package main

import (
	"gitee.com/xpigpig/vblog/cmd"
	"github.com/spf13/cobra"
)

func main() {
	err := cmd.Execute()
	cobra.CheckErr(err)
}
