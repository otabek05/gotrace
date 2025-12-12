package model

type Ethernet struct {
	SrcMAC       string  `json:"src_mac"`
	DstMAC       string  `json:"dst_mac"`
	EthernetType string  `json:"ethernet_type"`
	Length       uint16  `json:"length"`
}

type IPv4 struct {
	Version        int    `json:"version"`
	IHL            int    `json:"ihl"`
	TOS            uint8  `json:"tos"`
	TotalLength    uint16 `json:"total_length"`
	Identification uint16 `json:"identification"`
	Flags          string `json:"flags"`
	FragmentOffset uint16 `json:"fragment_offset"`
	TTL            uint8  `json:"ttl"`
	Protocol       string `json:"protocol"`
	Checksum       uint16 `json:"checksum"`

	SrcIP string `json:"src_ip"`
	DstIP string `json:"dst_ip"`
}

type TCP struct {
	SrcPort string `json:"src_port"`
	DstPort string  `json:"dst_port"`

	Seq uint32 `json:"seq"`
	Ack uint32 `json:"ack"`

	DataOffset uint8 `json:"data_offset"`
	Flags      []string `json:"flags"` 
	Window     uint16 `json:"window"`
	Checksum   uint16 `json:"checksum"`
	Urgent     uint16 `json:"urgent"`
	Options []string `json:"options"`
}

type UDP struct {
	SrcPort  uint16 `json:"src_port"`
	DstPort  uint16 `json:"dst_port"`
	Length   uint16 `json:"length"`
	Checksum uint16 `json:"checksum"`
}



type DNSQuestion struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Class string `json:"class"`
}


type DNSAnswer struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Class string `json:"class"`
	TTL uint32 `json:"ttl"`
	Data string `json:"data"` 
}


type DNS struct {
	ID      uint16        `json:"id"`
	QR      string        `json:"qr"` // Query / Response
	Opcode  string        `json:"opcode"`
	AA      bool          `json:"aa"`
	TC      bool          `json:"tc"`
	RD      bool          `json:"rd"`
	RA      bool          `json:"ra"`
	Z       uint8         `json:"z"`
	ResponseCode string   `json:"response_code"`

	Questions []DNSQuestion `json:"questions"`
	Answers   []DNSAnswer   `json:"answers"`
}



type IPv6 struct {
    TrafficClass uint8  `json:"traffic_class"`
    FlowLabel    uint32 `json:"flow_label"`
    PayloadLen   uint16 `json:"payload_length"`
    NextHeader   uint8  `json:"next_header"`
    HopLimit     uint8  `json:"hop_limit"`
    SrcIP        string `json:"src_ip"`
    DstIP        string `json:"dst_ip"`
}


type ICMP struct {
    Type     uint8  `json:"type"`
    Code     uint8  `json:"code"`
    Checksum uint16 `json:"checksum"`
}


type ARP struct {
    Operation string `json:"operation"`
    SenderMAC string `json:"sender_mac"`
    SenderIP  string `json:"sender_ip"`
    TargetMAC string `json:"target_mac"`
    TargetIP  string `json:"target_ip"`
}

type HTTP struct {
    Method      string            `json:"method"`
    URL         string            `json:"url"`
    Version     string            `json:"version"`
    StatusCode  int               `json:"status_code"`
    Status      string            `json:"status"`
    Headers     map[string]string `json:"headers"`
    Body        string            `json:"body"`
}



type HTTPS struct {
    TLSVersion string `json:"tls_version"`
    Cipher     string `json:"cipher"`
    ServerName string `json:"server_name"`
}



type DHCP struct {
    Operation      string `json:"operation"`
    TransactionID  string `json:"transaction_id"`
    ClientIP       string `json:"client_ip"`
    YourIP         string `json:"your_ip"`
    ServerIP       string `json:"server_ip"`
    GatewayIP      string `json:"gateway_ip"`
    Hostname       string `json:"hostname,omitempty"`
    LeaseTime      uint32 `json:"lease_time,omitempty"`
}


type AppLayer struct {
    Protocol string      `json:"protocol"`
    Data     interface{} `json:"data"`
}
