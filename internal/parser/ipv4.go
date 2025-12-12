package parser

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

func (p *PacketParser) ParseIPv4(layer *layers.IPv4, packet *model.ParsedPacket) {

	ipv4 := &model.IPv4{
		Version:        int(layer.Version),
		IHL:            int(layer.IHL),
		TOS:            layer.TOS,
		TotalLength:    layer.Length,
		Identification: layer.Id,
		Flags:          layer.Flags.String(),
		FragmentOffset: layer.FragOffset,
		TTL:            layer.TTL,
		Protocol:       layer.Protocol.String(),
		Checksum:       layer.Checksum,
		SrcIP:          layer.SrcIP.String(),
		DstIP:          layer.DstIP.String(),
	}

	packet.IPv4 = ipv4
}
