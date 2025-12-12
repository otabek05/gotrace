package model

type LayerInfo struct {
	Fields  map[string]string `json:"fields,omitempty"`
	Payload *LayerInfo        `json:"payload,omitempty"`
}


type ParsedPacket struct {
    Ethernet *Ethernet `json:"ethernet,omitempty"`
    IPv4     *IPv4     `json:"ipv4,omitempty"`
    IPv6     *IPv6     `json:"ipv6,omitempty"`
    TCP      *TCP      `json:"tcp,omitempty"`
    UDP      *UDP      `json:"udp,omitempty"`
    ICMP     *ICMP     `json:"icmp,omitempty"`
    ARP      *ARP      `json:"arp,omitempty"`
    DNS      *DNS      `json:"dns,omitempty"`
    HTTP     *HTTP     `json:"http,omitempty"`
    HTTPS    *HTTPS    `json:"https,omitempty"`
    DHCP     *DHCP     `json:"dhcp,omitempty"`
    App      *AppLayer `json:"application,omitempty"`
}
