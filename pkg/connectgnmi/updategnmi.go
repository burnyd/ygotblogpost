package connectgnmi

import (
	"context"
	"log"

	"github.com/aristanetworks/goarista/gnmi"
)

func Set(Target, Port, Username, Password, ntpserver string) {

	var cfg = &gnmi.Config{
		Addr:     Target + ":" + Port,
		Username: Username,
		Password: Password,
	}
	paths := []string{"/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server[address=" + ntpserver + "]/openconfig-system:config/address"}
	log.Print("\n trying to update with ntp server address: ", ntpserver)
	log.Print("\n Trying Path: ", paths)
	var origin = "openconfig"
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	var setOps []*gnmi.Operation

	op := &gnmi.Operation{
		Type:   "update",
		Path:   gnmi.SplitPath(paths[0]),
		Origin: origin,
		Val:    ntpserver,
	}

	setOps = append(setOps, op)

	err = gnmi.Set(ctx, client, setOps)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("ntp server configured: ", ntpserver)
	}
}
