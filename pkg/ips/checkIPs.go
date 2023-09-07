package ips

import (
	"log"
	"regexp"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	classAPrivate, classBPrivate, classCPrivate *regexp.Regexp
)

func init() {
	classAPrivate = regexp.MustCompile(`^10\.(([1-9]?\d|[12]\d\d)\.){2}([1-9]?\d|[12]\d\d)$`)
	classBPrivate = regexp.MustCompile(`^172\.(1[6-9]|2\d|3[01])(\.([1-9]?\d|[12]\d\d)){2}$`)
	classCPrivate = regexp.MustCompile(`^192\.16[6-8](\.([1-9]?\d|[12]\d\d)){2}$`)
}

// CheckIPClass checks the IP class read from the chosen interface
func ReadPacket(input string) {
	// 1.0 First get the packet
	if handle, err := pcap.OpenLive(input, 1600, true, pcap.BlockForever); err != nil {
		log.Println(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil { // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			ipPacket, _ := ipLayer.(*layers.IPv4)

			// 2.0 check ip class
			if matchIP(ipPacket.SrcIP.String()) == true {
				log.Println("======> Private IP read")
			} else {
				log.Println("======> Public IP read")
			}

		}
	}

}

func matchIP(ip string) bool {
	classAMatch := classAPrivate.MatchString(ip)
	classBMatch := classBPrivate.MatchString(ip)
	classCMatch := classCPrivate.MatchString(ip)

	if classAMatch == true || classBMatch == true || classCMatch == true {
		return true
	}

	return false
}
