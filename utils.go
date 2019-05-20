package main

import (
	"log"
	"net"
	"strconv"
)

const ipPath = "./ipipfree.ipdb"

func IpNToA(ip string) net.IP {
	if len(ip) >= 18 {
		runes := []rune(ip)
		ip = string(runes[0:18])
	}
	ipInt64, err := strconv.ParseInt(ip, 10, 64)
	if err != nil {
		log.Fatalf("convert %s failed", ip)
	}
	var bytes [4]byte
	bytes[0] = byte(ipInt64 & 0xFF)
	bytes[1] = byte((ipInt64 >> 8) & 0xFF)
	bytes[2] = byte((ipInt64 >> 16) & 0xFF)
	bytes[3] = byte((ipInt64 >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
