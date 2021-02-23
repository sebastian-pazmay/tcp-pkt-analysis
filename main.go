package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"tcp-pkt-analysis/utils"
)

var (
	filter       string = "ip proto tcp"
	pcapFilePath string = "capture/tcp.pcap"
	pcapHandle   *pcap.Handle
	err          error
)

func main() {
	pcapHandle, err = pcap.OpenOffline(pcapFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer pcapHandle.Close()
	err = pcapHandle.SetBPFFilter(filter)
	pktSrc := gopacket.NewPacketSource(pcapHandle, pcapHandle.LinkType())
	for packet := range pktSrc.Packets() {
		// fmt.Printf("%T\n", packet)
		utils.PayloadTCP(packet)
	}
}
