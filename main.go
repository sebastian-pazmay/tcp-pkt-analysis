package main

import (
    // "fmt"
    "log"
	"tcp-pkt-analysis/utils"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

var (
    pcapFilePath string = "capture/tcp.pcap"
    pcapFile   *pcap.Handle
    err      error
)

func main() {
    pcapFile, err = pcap.OpenOffline(pcapFilePath)
    if err != nil {
        log.Fatal(err)
    }
    defer pcapFile.Close()
    pktSrc := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())
    for packet := range pktSrc.Packets() {
        // fmt.Printf("%T\n", packet)
        go utils.PayloadTCP(packet)
    }
}
