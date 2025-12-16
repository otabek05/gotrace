package api

import (
	"net/http"
	"github.com/google/gopacket/pcap"
)

func NetInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	ifaces, err :=  pcap.FindAllDevs()
	if err != nil {
		BadResponse[string](w, err.Error())
		return 
	}

	var result []pcap.Interface

	for _, iface := range ifaces {
		hasIp := false

		for _, addr  := range iface.Addresses {
			if addr.IP != nil {
				hasIp = true
				break
			}
		}

		if !hasIp{
			continue
		}

		result = append(result, iface)
	}

	Success(w, result, "")
}