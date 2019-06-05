package utils

import (
	"log"
	"net"
	"strconv"
	"github.com/guoruibiao/ipservice/constants"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"github.com/pkg/errors"
)

func IpNToA(ip string) net.IP {
	if len(ip) >= 18 {
		runes := []rune(ip)
		ip = string(runes[0:18])
	}
	ipInt64, err := strconv.ParseInt(ip, 10, 64)
	if err != nil {
		log.Fatalf("convert %s failed", ip)
	}
	var bytes [4]byte
	bytes[0] = byte(ipInt64 & 0xFF)
	bytes[1] = byte((ipInt64 >> 8) & 0xFF)
	bytes[2] = byte((ipInt64 >> 16) & 0xFF)
	bytes[3] = byte((ipInt64 >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func GetOfficialData(ip string) (*constants.Entry, error) {
	if len(strings.Split(ip, "."))!=4 {
		return nil, errors.New("ip:" + ip + " format error. Should be like xx.xx.xx.xx")
	}
	resp, err := http.Get(constants.IPIP_OFFICIAL_URL + ip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var decoded []string
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return nil, err
	}
	entry := &constants.Entry{
		RegionName:decoded[1],
		CityName:decoded[2],
	}
	return entry, nil
}