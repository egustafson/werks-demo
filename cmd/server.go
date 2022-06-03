/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/egustafson/werks-demo/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run:   doServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func doServer(cmd *cobra.Command, args []string) {
	server.Start()
}
