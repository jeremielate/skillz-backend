package main

import (
	"log"
	"net"
)

func lookupHost(host string) string {
	hosts, err := net.LookupAddr(host)
	if err != nil {
		log.Println(err)
		return ""
	}
	return hosts[0]
}
