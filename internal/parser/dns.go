package parser

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

func (p *PacketParser) ParseDNS(layer *layers.DNS, packet *model.ParsedPacket) {
    dns := &model.DNS{
        ID:           layer.ID,
        QR:           dnsQRString(layer.QR),
        Opcode:       layer.OpCode.String(),
        AA:           layer.AA,
        TC:           layer.TC,
        RD:           layer.RD,
        RA:           layer.RA,
        Z:            layer.Z,
        ResponseCode: layer.ResponseCode.String(),
        Questions:    []model.DNSQuestion{},
        Answers:      []model.DNSAnswer{},
    }

    for _, q := range layer.Questions {
        dns.Questions = append(dns.Questions, model.DNSQuestion{
            Name:  string(q.Name),
            Type:  q.Type.String(),
            Class: q.Class.String(),
        })
    }

    for _, a := range layer.Answers {
        var data string

        switch a.Type {
        case layers.DNSTypeA, layers.DNSTypeAAAA:
            data = a.IP.String()
        case layers.DNSTypeCNAME:
            data = string(a.CNAME)
        case layers.DNSTypePTR:
            data = string(a.PTR)
        case layers.DNSTypeNS:
            data = string(a.NS)
        case layers.DNSTypeTXT:
            if len(a.TXTs) > 0 {
                data = string(a.TXTs[0])
            }
        default:
            data = string(a.Name)
        }

        dns.Answers = append(dns.Answers, model.DNSAnswer{
            Name:  string(a.Name),
            Type:  a.Type.String(),
            Class: a.Class.String(),
            TTL:   a.TTL,
            Data:  data,
        })
    }

    packet.DNS = dns
}

func dnsQRString(qr bool) string {
    if qr {
        return "response"
    }
    return "query"
}
