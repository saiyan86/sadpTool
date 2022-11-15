package main

type dahuaProbeReq struct {
	Method string `json:"method"`
	Params `json:"params"`
}
type Params struct {
	Mac string `json:"mac"`
	Uni int    `json:"uni"`
}

type dahuaProbeResp struct {
	Mac    string `json:"mac"`
	Method string `json:"method"`
	Params struct {
		DeviceInfo struct {
			AbroadInfo          string `json:"AbroadInfo"`
			AlarmInputChannels  int    `json:"AlarmInputChannels"`
			AlarmOutputChannels int    `json:"AlarmOutputChannels"`
			DeviceClass         string `json:"DeviceClass"`
			DeviceID            string `json:"DeviceID"`
			DeviceType          string `json:"DeviceType"`
			Find                string `json:"Find"`
			FindVersion         int    `json:"FindVersion"`
			HTTPPort            int    `json:"HttpPort"`
			IPv4Address         struct {
				DefaultGateway string `json:"DefaultGateway"`
				DhcpEnable     bool   `json:"DhcpEnable"`
				IPAddress      string `json:"IPAddress"`
				SubnetMask     string `json:"SubnetMask"`
			} `json:"IPv4Address"`
			IPv6Address struct {
				DefaultGateway   interface{} `json:"DefaultGateway"`
				DhcpEnable       bool        `json:"DhcpEnable"`
				IPAddress        string      `json:"IPAddress"`
				LinkLocalAddress string      `json:"LinkLocalAddress"`
			} `json:"IPv6Address"`
			Init                     int    `json:"Init"`
			MachineGroup             string `json:"MachineGroup"`
			MachineName              string `json:"MachineName"`
			Manufacturer             string `json:"Manufacturer"`
			Port                     int    `json:"Port"`
			RemoteVideoInputChannels int    `json:"RemoteVideoInputChannels"`
			SerialNo                 string `json:"SerialNo"`
			UnLoginFuncMask          int    `json:"UnLoginFuncMask"`
			Vendor                   string `json:"Vendor"`
			Version                  string `json:"Version"`
			VideoInputChannels       int    `json:"VideoInputChannels"`
			VideoOutputChannels      int    `json:"VideoOutputChannels"`
		} `json:"deviceInfo"`
	} `json:"params"`
}

const (
	// for now we use request hex directly, for we don't know how construct of this request
	daHuaReq = "20000000444849500000000000000000490000000000000049000000000000007b20226d6574686f6422203a20224448446973636f7665722e736561726368222c2022706172616d7322203a207b20226d616322203a2022222c2022756e6922203a2031207d207d0a"
)
