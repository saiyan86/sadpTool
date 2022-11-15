package main

import "encoding/xml"

type Probe struct {
	Uuid  string `xml:"Uuid"`
	Types string `xml:"Types"`
}

// generated: https://www.onlinetool.io/xmltogo/
type HikvProbeResp struct {
	XMLName                  xml.Name `xml:"ProbeMatch"`
	Uuid                     string   `xml:"Uuid"`
	Types                    string   `xml:"Types"`
	DeviceType               string   `xml:"DeviceType"`
	DeviceDescription        string   `xml:"DeviceDescription"`
	DeviceSN                 string   `xml:"DeviceSN"`
	CommandPort              string   `xml:"CommandPort"`
	HttpPort                 string   `xml:"HttpPort"`
	MAC                      string   `xml:"MAC"`
	IPv4Address              string   `xml:"IPv4Address"`
	IPv4SubnetMask           string   `xml:"IPv4SubnetMask"`
	IPv4Gateway              string   `xml:"IPv4Gateway"`
	IPv6Address              string   `xml:"IPv6Address"`
	IPv6Gateway              string   `xml:"IPv6Gateway"`
	IPv6MaskLen              string   `xml:"IPv6MaskLen"`
	DHCP                     string   `xml:"DHCP"`
	AnalogChannelNum         string   `xml:"AnalogChannelNum"`
	DigitalChannelNum        string   `xml:"DigitalChannelNum"`
	SoftwareVersion          string   `xml:"SoftwareVersion"`
	DSPVersion               string   `xml:"DSPVersion"`
	BootTime                 string   `xml:"BootTime"`
	Encrypt                  string   `xml:"Encrypt"`
	ResetAbility             string   `xml:"ResetAbility"`
	DiskNumber               string   `xml:"DiskNumber"`
	Activated                string   `xml:"Activated"`
	PasswordResetAbility     string   `xml:"PasswordResetAbility"`
	PasswordResetModeSecond  string   `xml:"PasswordResetModeSecond"`
	SupportSecurityQuestion  string   `xml:"SupportSecurityQuestion"`
	SupportHCPlatform        string   `xml:"SupportHCPlatform"`
	HCPlatformEnable         string   `xml:"HCPlatformEnable"`
	IsModifyVerificationCode string   `xml:"IsModifyVerificationCode"`
	Salt                     string   `xml:"Salt"`
	DeviceLock               string   `xml:"DeviceLock"`
	SupportMailBox           string   `xml:"SupportMailBox"`
	WifiMAC                  string   `xml:"WifiMAC"`
	SupportEzvizUnbind       string   `xml:"supportEzvizUnbind"`
}

const (
	HikvisionPort = ""
)
