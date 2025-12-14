package model

type WSReceiveMessage struct {
	Type                string   `json:"type"`
	Message             any      `json:"message"`
	TrafficOptions      string   `json:"trafficOptions"`
	NetworkLayer        string   `json:"networkLayer"`
	Transport           string   `json:"transport"`
	ApplicationServices []string `json:"services"`
}

type WSSendingMessage struct {
	Type    string        `json:"type"`
	Packets *ParsedPacket `json:"packets,omitempty"`
}
