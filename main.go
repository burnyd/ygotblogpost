package main

import (
	"flag"
	"fmt"

	"github.com/burnyd/ygotblogpost/pkg/connectgnmi"
	"github.com/burnyd/ygotblogpost/pkg/createjson"
	"github.com/openconfig/ygot/ygot"
)

func main() {
	Target := flag.String("target", "172.20.20.2", "gnmi target")
	Port := flag.String("port", "6030", "gNMI port default is 6030")
	Username := flag.String("username", "admin", "admin")
	Password := flag.String("password", "admin", "admin")
	NtpServerAddress := flag.String("ntpserveraddress", "", "Address in which you want to render a NtpServerAddress")
	GetNtPServers := flag.Bool("getntpservers", false, "Uses the gNMI GET Method to get all the NTP Servers")
	SetNtpAddress := flag.String("setntpaddress", "", "Address in which you want to set as the ntp server.")
	DeleteNtpAddress := flag.String("deletentpaddress", "", "ntp server you want to delete")
	Subscribe := flag.Bool("subscribe", false, "Subscribe method to the ntp servers")
	flag.Parse()
	if *NtpServerAddress != "" {
		createjson.CreateNtpJson(ygot.String(*NtpServerAddress))
	}
	if *SetNtpAddress != "" {
		connectgnmi.Set(*Target, *Port, *Username, *Password, *SetNtpAddress)
	}
	if *DeleteNtpAddress != "" {
		connectgnmi.Delete(*Target, *Port, *Username, *Password, *DeleteNtpAddress)
	}
	if *GetNtPServers == true {
		connectgnmi.Get(*Target, *Port, *Username, *Password)
	}
	if *Subscribe == true {
		connectgnmi.Subscribe(*Target, *Port, *Username, *Password)
	}
	if flag.NFlag() == 0 {
		fmt.Println("You need to enter a thing please!")
	}
	//} else {
	//	connectgnmi.Get(*Target, *Port, *Username, *Password)
	//}
	// Create a type for ntpserver.  Typically, we would pass this into a function just for demo.  Give it an IPv4 address.
	/*var ntpserver *string
	ntpserverIp := "1.2.3.4"
	ntpserver = &ntpserverIp

	NtpServer := ocsystem.System_Ntp_Server{
		Address: ygot.String(*ntpserver),
	}

	NtpMap := make(map[string]*ocsystem.System_Ntp_Server)

	NtpMap[*ntpserver] = &NtpServer

	NtpSys := ocsystem.System_Ntp{
		Server: NtpMap,
	}

	Sys := &ocsystem.System{Ntp: &NtpSys}

	json, err := ygot.EmitJSON(Sys, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Value error: %v", err))
	}
	// Print the json
	fmt.Println(json)
	*/

}
