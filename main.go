package main

import (
	"context"
	"log"

	"github.com/comeonjy/account/cmd"
	"github.com/comeonjy/account/pkg/consts"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(c *cobra.Command, args []string) {
		xenv.Init(consts.EnvMap)
		logger := xlog.New(xlog.WithTrace(xenv.GetEnv(xenv.TraceName)))
		ctx, cancel := context.WithCancel(context.Background())
		app := cmd.InitApp(ctx, logger)
		if err := app.Run(cancel); err != nil {
			log.Println("Server Exit:", err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
