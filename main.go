package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"tcp-pkt-analysis/utils"
)

var (
	filter       string = "vlan and tcp"
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
	if err != nil {
		log.Fatal(err)
	}
	pktSrc := gopacket.NewPacketSource(pcapHandle, pcapHandle.LinkType())
	counter := 0
	for packet := range pktSrc.Packets() {
		counter++
		// fmt.Printf("%T\n", packet)
		utils.PayloadPacket(packet)
	}
	fmt.Printf("Total filtered pkts in pcap: %v", counter)
}
