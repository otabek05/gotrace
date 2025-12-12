package parser

import (
	"gotrace/internal/model"
	"strconv"
	"strings"
)

func (p *PacketParser) ParseHTTP(payload []byte, packet *model.ParsedPacket) {
	data := string(payload)
	http := &model.HTTP{
		Headers: make(map[string]string),
	}

	lines := strings.Split(data, "\r\n")

	if len(lines) == 0 {
		return
	}

	startLine := lines[0]

	if strings.HasPrefix(startLine, "GET ") ||
		strings.HasPrefix(startLine, "POST ") ||
		strings.HasPrefix(startLine, "PUT ") ||
		strings.HasPrefix(startLine, "DELETE ") ||
		strings.HasPrefix(startLine, "HEAD ") ||
		strings.HasPrefix(startLine, "OPTIONS ") {

		parts := strings.Split(startLine, " ")
		if len(parts) >= 3 {
			http.Method = parts[0]
			http.URL = parts[1]
			http.Version = parts[2]
		}

	} else if strings.HasPrefix(startLine, "HTTP/") {

		parts := strings.SplitN(startLine, " ", 3)
		if len(parts) >= 3 {
			http.Version = parts[0]
			code, _ := strconv.Atoi(parts[1])
			http.StatusCode = code
			http.Status = parts[2]
		}
	}


	i := 1
	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		headerParts := strings.SplitN(line, ": ", 2)
		if len(headerParts) == 2 {
			http.Headers[headerParts[0]] = headerParts[1]
		}
	}

	if i+1 < len(lines) {
		http.Body = strings.Join(lines[i+1:], "\n")
	}

	packet.HTTP = http

}
