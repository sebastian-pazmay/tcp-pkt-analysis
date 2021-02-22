package utils

import (
    "fmt"
)

func PayloadTCP(pkt *gopacket.eagerPacket) {
	fmt.Println("PACKET:")
	fmt.Println(packet)
}
