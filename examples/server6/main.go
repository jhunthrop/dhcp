package main

import (
	"log"
	"net"

	"github.com/andrewrynhard/dhcp/dhcpv6"
	"github.com/andrewrynhard/dhcp/dhcpv6/server6"
)

func handler(conn net.PacketConn, peer net.Addr, m dhcpv6.DHCPv6) {
	// this function will just print the received DHCPv6 message, without replying
	log.Print(m.Summary())
}

func main() {
	laddr := net.UDPAddr{
		IP:   net.ParseIP("::1"),
		Port: dhcpv6.DefaultServerPort,
	}
	server := server6.NewServer(laddr, handler)

	defer server.Close()
	if err := server.ActivateAndServe(); err != nil {
		log.Panic(err)
	}
}
