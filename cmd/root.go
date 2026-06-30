/*
Copyright © 2026 Masato Kikuchi <m-kikuchi@neody.ad.jp>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "etherip-client",
	Short: "Ethernet over IPv6(RFC3378)カーネルのクライアント",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


