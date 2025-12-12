package parser

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

func (e *PacketParser) ParseEthernet(eth *layers.Ethernet, packet *model.ParsedPacket) {

	ethLayer := &model.Ethernet{
		SrcMAC:       eth.SrcMAC.String(),
		DstMAC:       eth.DstMAC.String(),
		EthernetType: eth.EthernetType.String(),
		Length:       eth.Length,
	}

	packet.Ethernet = ethLayer

}
