//go:build darwin
// +build darwin

package gokb

import (
	"net"
	"time"
)

func CreateDialer(timeout timeoutParams) net.Dialer {
	return net.Dialer{
		Timeout:   time.Duration(timeout.connect_timeout) * time.Second,
		KeepAlive: time.Duration(timeout.keepalive_interval) * time.Second,
	}
}
