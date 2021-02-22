package utils

import (
    "fmt"
	"github.com/google/gopacket"
)

func PayloadTCP(pkt gopacket.Packet) {
	fmt.Println(pkt)
	applicationLayer := pkt.ApplicationLayer()
	fmt.Printf("%s\n", applicationLayer.Payload())
	fmt.Printf("%T\n", applicationLayer.Payload())
}
