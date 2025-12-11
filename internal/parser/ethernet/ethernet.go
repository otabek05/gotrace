package ethernet

import (
	"gotrace/internal/model"

	"github.com/google/gopacket/layers"
)

type EthernetParser struct{}

func New() *EthernetParser { return &EthernetParser{} }

func (e *EthernetParser) Name() string { return "Ethernet" }

func (e *EthernetParser) Parse(eth *layers.Ethernet) *model.LayerInfo {

	info := model.LayerInfo{
		Name: e.Name(),
		Fields: map[string]string{
			"Source":eth.SrcMAC.String(),
			"Destination": eth.DstMAC.String(),
			"Type": eth.EthernetType.String(),
		},
	}

	return &info

}
