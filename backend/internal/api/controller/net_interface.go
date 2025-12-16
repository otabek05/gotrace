package controller

import (
	"net"
	"github.com/gin-gonic/gin"
	"github.com/google/gopacket/pcap"
)

func (h *handler) getNetInterfaces(c *gin.Context) {
	ifaces, err :=  pcap.FindAllDevs()
	if err != nil {
		h.response.BadRequest(c, err.Error());
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

	h.response.Success(c, result, "") 
}


func (h *handler) getDomainIP(c *gin.Context) {
	domain := c.Query("domain")
	if domain == "" {
		h.response.NotFound(c, "no domain provided")
		return 
	}

	ips, err := net.LookupIP(domain)
	if err != nil {
		h.response.BadRequest(c, err.Error())
		return 
	}

	h.response.Success(c, ips, "success")
}