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

	Success(w, ifaces, "")

}