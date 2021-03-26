package uuid

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"net"
	"os"
)

func isPrivateIPv4(ip net.IP) bool {
	if ip == nil {
		return false
	}
	if ip[0] == 10 {
		return true
	}
	if ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) {
		return true
	}
	if ip[0] == 192 && ip[1] == 168 {
		return true
	}
	return false
}

func PrivateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}
		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, fmt.Errorf("no private ip address")
}

func CurrentPID() uint32 {
	return uint32(os.Getpid())
}

func GenHashCode(bytes []byte) uint16 {
	hash := fnv.New64a()
	hash.Write(bytes)
	binary.Write(hash, binary.LittleEndian, CurrentPID())
	return uint16(hash.Sum64())
}
