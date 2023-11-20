package main

import (
	"fmt"

	"github.com/burnyd/ygotblogpost/pkg/ocsystem"
	"github.com/openconfig/ygot/ygot"
)

func main() {
	// Create a type for ntpserver.  Typically, we would pass this into a function just for demo.  Give it an IPv4 address.
	var ntpserver *string
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
}
