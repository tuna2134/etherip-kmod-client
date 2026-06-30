/*
Copyright © 2026 Masato Kikuchi <m-kikuchi@neody.ad.jp>
*/
package cmd

import (
	"git.neody.ad.jp/tuna2134/etherip-client/utils"
	"github.com/spf13/cobra"
)

var DelDevVar string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete iface",
	Long: `Delete iface`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.DelDev(DelDevVar)
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&DelDevVar, "devname", "d", "eip0", "削除対象のNICの名前")
	rootCmd.AddCommand(deleteCmd)
}
