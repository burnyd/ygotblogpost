package createjson

import (
	"fmt"

	"github.com/burnyd/ygotblogpost/pkg/ocsystem"
	"github.com/openconfig/ygot/ygot"
)

func CreateNtpJson(server *string) string {
	NtpServer := ocsystem.System_Ntp_Server{
		Address: ygot.String(*server),
	}

	NtpMap := make(map[string]*ocsystem.System_Ntp_Server)

	NtpMap[*server] = &NtpServer

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
	return json
}