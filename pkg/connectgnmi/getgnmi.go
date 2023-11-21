package connectgnmi

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aristanetworks/goarista/gnmi"
	"github.com/burnyd/ygotblogpost/pkg/ocsystem"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func GetReq(ctx context.Context, client pb.GNMIClient,
	req *pb.GetRequest) ([]byte, error) {
	resp, err := client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	var reqreturn []string
	for _, notif := range resp.Notification {
		for _, update := range notif.Update {
			reqreturn = append(reqreturn, (gnmi.StrUpdateVal(update)))
		}
	}
	byteSlice := []byte(strings.Join(reqreturn, " "))
	return byteSlice, nil
}

func Get(Target, Port, Username, Password string) {
	var cfg = &gnmi.Config{
		Addr:     Target + ":" + Port,
		Username: Username,
		Password: Password,
	}
	paths := []string{"/openconfig-system:system/ntp/servers/server/config[address=*]"}
	var origin = "openconfig"
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	req, err := gnmi.NewGetRequest(gnmi.SplitPaths(paths), origin)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Addr != "" {
		if req.Prefix == nil {
			req.Prefix = &pb.Path{}
		}
		req.Prefix.Target = cfg.Addr
	}

	Get, err := GetReq(ctx, client, req)
	if err != nil {
		log.Fatal(err)
	}
	if len(Get) > 0 {
		fmt.Print("This is a string version of the servers \n")
		fmt.Print(string(Get), "\n")
		UnmarshallJson(Get)
	} else {
		fmt.Println("Get request did not receive any ntp servers")
	}
}

func UnmarshallJson(data []byte) {
	ReturnJson := &ocsystem.System_Ntp_Server{}

	if err := ocsystem.Unmarshal(data, ReturnJson); err != nil {
		panic(fmt.Sprintf("Cannot unmarshal JSON: %v", err))
	}
	fmt.Print("\n")
	fmt.Print("This is a Unmarshalled version of the data \n")
	fmt.Print(string(*ReturnJson.Address))
}
