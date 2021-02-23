package utils

import (
	"fmt"
	"github.com/google/gopacket"
)

func PayloadTCP(pkt gopacket.Packet) {
	applicationLayer := pkt.ApplicationLayer()
	pktPayload := applicationLayer.Payload()
	pktPayloadString := string(pktPayload[:])
	fmt.Printf("VALUE: %s\n", pktPayloadString)
	fmt.Printf("TYPE: %T\n", pktPayloadString)
}
