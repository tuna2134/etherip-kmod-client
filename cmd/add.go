/*
Copyright © 2026 Masato Kikuchi <m-kikuchi@neody.ad.jp>
*/
package cmd

import (
	"net"

	"git.neody.ad.jp/tuna2134/etherip-client/utils"
	"github.com/spf13/cobra"
)

var addDevName string
var addLocalIP net.IP
var addRemoteIP net.IP
var addLinkName string

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Ethernet over IPv6(RFC3378)カーネルのクライアント",
	RunE: AddrunE,
}

func AddrunE(cmd *cobra.Command, args []string) error {
	err := utils.AddDev(addDevName, addLocalIP, addRemoteIP, addLinkName)
	return err
}

func init() {
	AddCmd.Flags().StringVarP(&addDevName, "devname", "d", "eip0", "作成するNICの名")
	AddCmd.Flags().IPVarP(&addLocalIP, "local", "l", nil, "自分の終端アドレス")
	AddCmd.Flags().IPVarP(&addRemoteIP, "remote", "r", nil, "相手の終端アドレス")
	AddCmd.Flags().StringVarP(&addLinkName, "link", "L", "", "アンダーレイのNIC名")

	rootCmd.AddCommand(AddCmd)
}


