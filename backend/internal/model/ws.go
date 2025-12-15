package model

import "github.com/google/gopacket/pcap"

type WSReceiveMessage struct {
	Type                string         `json:"type"`
	Message             any            `json:"message"`
	TrafficOptions      string         `json:"trafficOptions"`
	NetworkLayer        string         `json:"networkLayer"`
	Transport           string         `json:"transport"`
	ApplicationServices []string       `json:"services"`
	NetworkInterface    pcap.Interface `json:"interface"`
}

type WSSendingMessage struct {
	Type    string        `json:"type"`
	Packets *ParsedPacket `json:"packets,omitempty"`
	InternetSpeed *InternetSpeed `json:"internetSpeed"`
}


type InternetSpeed struct {
	BytesIn string `json:"bytesIn"`
	BytesOut  string `json:"bytesOut"`
}