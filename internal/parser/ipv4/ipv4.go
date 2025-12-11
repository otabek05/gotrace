package ipv4

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

type IPV4Parser struct {}

func Net() *IPV4Parser {return &IPV4Parser{}}

func (i *IPV4Parser) Name() string {return "IPV4"}

func (i *IPV4Parser) Parse(ipv4 *layers.IPv4) *model.LayerInfo {

    info := model.LayerInfo{
        Name: "IPv4",
        Fields: map[string]string{
            "Source":   ipv4.SrcIP.String(),
            "Destination": ipv4.DstIP.String(),
            "Protocol": ipv4.Protocol.String(),
        },
    }

    return &info

}
