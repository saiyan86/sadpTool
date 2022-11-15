package main

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	Header = `<?xml version="1.0" encoding="UTF-8"?>`
)

func main() {
	go getDahua()
	go getHikv()
	select {}
}
func getDahua() {
	lAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 37809,
	}
	rAddr := &net.UDPAddr{
		IP:   net.IPv4(239, 255, 255, 251),
		Port: 37810,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Println(err)
	}
	go func() {
		result := &dahuaProbeResp{}
		for {
			buffer := make([]byte, 2048)
			n, _, _ := conn.ReadFromUDP(buffer)
			if n == 0 {
				return
			}
			buffer = buffer[32 : n-2]
			err = json.Unmarshal(buffer, &result)
			if err != nil {
				log.Println("err to unmarshal response to json ,error: ", err.Error())
				continue
			}
			log.Println("======= Searched a Dahua Device! =========")
			log.Println("ip:", result.Params.DeviceInfo.IPv4Address.IPAddress)
			log.Println("s/n:", result.Params.DeviceInfo.SerialNo)
			log.Println("==========================================")
		}
	}()
	defer conn.Close()

	marshal, err := hex.DecodeString(daHuaReq)
	if err != nil {
		panic(err)
	}
	sendData := marshal
	_, err = conn.WriteTo(sendData, rAddr)
	if err != nil {
		fmt.Println("Search Request Send Failed err:", err)
	}
	time.Sleep(time.Second)
}

func getHikv() {
	lAddr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 37020,
	}
	rAddr := &net.UDPAddr{
		IP:   net.IPv4(239, 255, 255, 250),
		Port: 37020,
	}
	conn, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		log.Println(err)
	}
	go func() {
		buffer := make([]byte, 2048)
		result := &HikvProbeResp{}
		for {
			_, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				log.Println(err.Error())
				return
			}
			err = xml.Unmarshal(buffer, &result)
			if err != nil {
				log.Println("err to unmarshal response to xml ,error: ", err.Error())
				continue
			}
			log.Println("==== Searched a Hikvision Device! ========")
			log.Println("IP:", result.IPv4Address)
			log.Println("序列号", result.DeviceSN[len(result.DeviceSN)-9:])
			log.Println("==========================================")
		}
	}()
	defer conn.Close()

	uid := strings.ToUpper(uuid.NewV4().String())
	req := Probe{
		Uuid:  uid,
		Types: "inquiry",
	}
	marshal, err := xml.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	marshal = append([]byte(Header), marshal...)

	sendData := marshal
	_, err = conn.WriteTo(sendData, rAddr)
	if err != nil {
		fmt.Println("Search Request Send Failed err:", err)
	}
	time.Sleep(time.Second)
}
