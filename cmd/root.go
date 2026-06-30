/*
Copyright © 2026 Masato Kikuchi <m-kikuchi@neody.ad.jp>
*/
package cmd

import (
	"net"
	"os"

	"git.neody.ad.jp/tuna2134/etherip-client/utils"
	"github.com/spf13/cobra"
)

var devName string
var localIP net.IP
var remoteIP net.IP
var linkName string

var rootCmd = &cobra.Command{
	Use:   "etherip-client",
	Short: "Ethernet over IPv6(RFC3378)カーネルのクライアント",
	RunE: runE,
}

func runE(cmd *cobra.Command, args []string) error {
	err := utils.AddDev(devName, localIP, remoteIP, linkName)
	return err
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&devName, "devname", "d", "eip0", "作成するNICの名")
	rootCmd.Flags().IPVarP(&localIP, "local", "l", nil, "自分の終端アドレス")
	rootCmd.Flags().IPVarP(&remoteIP, "remote", "r", nil, "相手の終端アドレス")
	rootCmd.Flags().StringVarP(&linkName, "link", "L", "", "アンダーレイのNIC名")
}


