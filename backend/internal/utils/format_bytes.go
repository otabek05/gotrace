package utils

import (
	"encoding/hex"
	"fmt"
)


func FormatBytesPerSec(b uint64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case b >= GB:
		return fmt.Sprintf("%.2f GB/s", float64(b)/GB)
	case b >= MB:
		return fmt.Sprintf("%.2f MB/s", float64(b)/MB)
	case b >= KB:
		return fmt.Sprintf("%.2f KB/s", float64(b)/KB)
	default:
		return fmt.Sprintf("%d B/s", b)
	}
}


func BytesToSafeString(b []byte) string {
	for _, c := range b {
		if c < 9 || (c > 13 && c < 32) {
			return hex.Dump(b)
		}
	}
	return string(b)
}
