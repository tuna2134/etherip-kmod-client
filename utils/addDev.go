/*
Copyright © 2026 Masato Kikuchi <m-kikuchi@neody.ad.jp>
*/

package utils

import (
	"encoding/binary"
	"net"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"golang.org/x/sys/unix"
)

const (
	IFLA_ETHERIP6_UNSPEC = iota
	IFLA_ETHERIP6_LOCAL
	IFLA_ETHERIP6_REMOTE
	IFLA_ETHERIP6_LINK
	IFLA_ETHERIP6_HOP_LIMIT
)

func u32(v uint32) []byte {
	b := make([]byte, 4)
	binary.NativeEndian.PutUint32(b, v)
	return b
}

func AddDev(DevName string, local net.IP, remote net.IP, linkName string) error {
	req := nl.NewNetlinkRequest(
		unix.RTM_NEWLINK,
		unix.NLM_F_REQUEST|unix.NLM_F_CREATE|unix.NLM_F_EXCL|unix.NLM_F_ACK,
	)

	msg := nl.NewIfInfomsg(unix.AF_UNSPEC)
	req.AddData(msg)

	req.AddData(nl.NewRtAttr(unix.IFLA_IFNAME, []byte(DevName)))

	linkInfo := nl.NewRtAttr(unix.IFLA_LINKINFO, nil)
	linkInfo.AddRtAttr(nl.IFLA_INFO_KIND, nl.ZeroTerminated("etherip6"))

	infoData := nl.NewRtAttr(nl.IFLA_INFO_DATA, nil)
	infoData.AddRtAttr(IFLA_ETHERIP6_LOCAL, local.To16())
	infoData.AddRtAttr(IFLA_ETHERIP6_REMOTE, remote.To16())

	if linkName != "" {
		link, err := netlink.LinkByName(linkName)
		if err != nil {
			return err
		}
		infoData.AddRtAttr(IFLA_ETHERIP6_LINK, u32(uint32(link.Attrs().Index)))
	}

	infoData.AddRtAttr(IFLA_ETHERIP6_HOP_LIMIT, []byte{255})

	linkInfo.AddRtAttr(nl.IFLA_INFO_DATA, infoData.Serialize())
	req.AddData(linkInfo)

	_, err := req.Execute(unix.NETLINK_ROUTE, 0)

	if err != nil {
		return err
	}
	return nil
}