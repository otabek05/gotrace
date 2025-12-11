package http

import (

	"gotrace/internal/model"

	"github.com/google/gopacket"
)

type HTTPParser struct{}

func New() *HTTPParser { return &HTTPParser{} }

func (p *HTTPParser) Name() string { return "HTTP" }

func (p *HTTPParser) Parse(packet *gopacket.Packet) *model.LayerInfo {
    
    info := model.LayerInfo{
        Name:   "HTTP",
        Fields: map[string]string{
            "Payload": "",
        },
    }

    return &info
}
