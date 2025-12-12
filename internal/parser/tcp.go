package parser

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

func (p *PacketParser) ParseTCP(tcpLayer *layers.TCP, packet *model.ParsedPacket) {

	flags := []string{}
	if tcpLayer.SYN {
		flags = append(flags, "SYN")
	}
	if tcpLayer.ACK {
		flags = append(flags, "ACK")
	}
	if tcpLayer.FIN {
		flags = append(flags, "FIN")
	}
	if tcpLayer.RST {
		flags = append(flags, "RST")
	}
	if tcpLayer.PSH {
		flags = append(flags, "PSH")
	}
	if tcpLayer.URG {
		flags = append(flags, "URG")
	}
	if tcpLayer.ECE {
		flags = append(flags, "ECE")
	}
	if tcpLayer.CWR {
		flags = append(flags, "CWR")
	}

	tcp := &model.TCP{
		SrcPort:    tcpLayer.SrcPort.String(),
		DstPort:    tcpLayer.DstPort.String(),
		Seq:        tcpLayer.Seq,
		Ack:        tcpLayer.Ack,
		DataOffset: tcpLayer.DataOffset,
		Window:     tcpLayer.Window,
		Checksum:   tcpLayer.Checksum,
		Urgent:     tcpLayer.Urgent,
		Flags:      flags,
	}

	packet.TCP = tcp
}
