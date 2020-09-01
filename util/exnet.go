package util

import (
	"net"
	"net/http"
	"strings"
)

func ClientIP(r *http.Request) string {
	// 尝试从 X-Forwarded-For 中获取
	xForwardedFor := r.Header.Get(`X-Forwarded-For`)
	ip := strings.TrimSpace(strings.Split(xForwardedFor, `,`)[0])
	if ip == `` {
		// 尝试从 X-Real-Ip 中获取
		ip = strings.TrimSpace(r.Header.Get(`X-Real-Ip`))
		if ip == `` {
			// 直接从 Remote Addr 中获取
			_ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
			if err != nil {
				panic(err)
			} else {
				ip = _ip
			}
		}
	}
	// 从控制台输出获取到的IP地址
	return ip
}

func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}