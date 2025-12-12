package parser

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)


func (p *PacketParser) ParseUDP(layer *layers.UDP, packet *model.ParsedPacket) {
	 udpPacket := &model.UDP{  
        SrcPort: uint16(layer.SrcPort),
        DstPort: uint16(layer.DstPort),
        Length: layer.Length,
        Checksum: layer.Checksum,
    }

	packet.UDP = udpPacket
}