// Version = "0.0.1"

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func cidrToIPs(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}

	ips := make([]string, 0)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		for _, ip := range cidrToIPs(sc.Text()) {
			fmt.Println(ip)
		}
	}
}
