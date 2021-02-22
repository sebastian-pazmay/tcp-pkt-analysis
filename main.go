package main

import (
    "log"
	"tcp-pkt-analysis/utils"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)


var (
    pcapFilePath string = "capture/tcp.pcap"
    pcapFile     *pcap.Handle
    err          error
    wg           sync.WaitGroup
)

func main() {
    
    wg.Add(sliceLength)

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
    wg.Wait()
}


https://medium.com/@greenraccoon23/multi-thread-for-loops-easily-and-safely-in-go-a2e915302f8b
http://www.golangpatterns.info/concurrency/parallel-for-loop