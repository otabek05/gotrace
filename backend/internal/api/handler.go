package api

import (
	"gotrace/backend/internal/capture"
	"gotrace/backend/internal/types"
	"gotrace/backend/internal/ws"
	"net"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex

func listenInterfaces(c *gin.Context) {
	ifs, err := net.Interfaces()
	if err != nil {
	     error(c, err.Error(), 500)
		return
	}

	out := make([]types.NetworkInterface, 0, len(ifs))
	for _, ifi := range ifs {
		addrs,  _ := ifi.Addrs()
		as := make([]string, 0, len(addrs))
		for _, a := range addrs {
			as = append(as, a.String())
		}

		out = append(out, types.NetworkInterface{
			Name:         ifi.Name,
			MTU:          ifi.MTU,
			HardwareAddr: ifi.HardwareAddr.String(),
			Addresses:    as,
			IsUp:         (ifi.Flags&net.FlagUp != 0),
		})
	}

	c.JSON(200, types.ApiResponse[any]{
		Success: true,
		Message: "",
		Data: out,
	})

}

type startReq struct {
	Interface string `json:"interface"`
	Snaplen   int32  `json:"snaplen"`
	Promisc   bool   `json:"promisc"`
	Filter    string `json:"filter"`
}


func StartScan(c *gin.Context) {
	var req startReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, types.ApiResponse[any]{
			Success: false,
			Message: "invalid body: " + err.Error(),
		})
		return
	}

	if req.Snaplen == 0 {
		req.Snaplen = 65535
	}

	mu.Lock()
	defer mu.Unlock()

	if running, _ := capture.DefaultScanner.IsRunning(); running {
		c.JSON(409, types.ApiResponse[any]{
			Success: false,
			Message: "scanner already running",
		})
		return
	}

	if err := capture.DefaultScanner.Start(
		req.Interface,
		req.Snaplen,
		req.Promisc,
		req.Filter,
	); err != nil {
		c.JSON(500, types.ApiResponse[any]{
			Success: false,
			Message: "failed to start: " + err.Error(),
		})
		return
	}

	c.JSON(200, types.ApiResponse[any]{
		Success: true,
		Message: "started",
	})
}


func StopScan(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	capture.DefaultScanner.Stop()

	c.JSON(200, types.ApiResponse[any]{
		Success: true,
		Message: "stopped",
	})
}

func Status(c *gin.Context) {
	running, iface := capture.DefaultScanner.IsRunning()

	c.JSON(200, types.ApiResponse[any]{
		Success: true,
		Message: "status",
		Data: map[string]any{
			"running":   running,
			"interface": iface,
		},
	})
}

// ✅ GET /ws
func ServeWS(c *gin.Context) {
	ws.DefaultHub.ServeWS(c.Writer, c.Request)
}