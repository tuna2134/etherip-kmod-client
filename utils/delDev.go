package utils

import "github.com/vishvananda/netlink"

func DelDev(devname string) error {
	link, err := netlink.LinkByName(devname)
	if err != nil {
		return err
	}
	netlink.LinkDel(link)
	return nil
}