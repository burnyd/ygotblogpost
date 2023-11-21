package connectgnmi

import (
	"context"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/aristanetworks/goarista/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func Subscribe(Target, Port, Username, Password string) {
	var cfg = &gnmi.Config{
		Addr:     Target + ":" + Port,
		Username: Username,
		Password: Password,
	}

	paths := []string{"/openconfig-system:system/ntp/openconfig-system:servers/openconfig-system:server"}
	var origin = "openconfig"
	ctx := gnmi.NewContext(context.Background(), cfg)
	client, err := gnmi.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	subOptions := gnmi.SubscribeOptions{
		Origin: origin,
		Paths:  gnmi.SplitPaths(paths),
		Target: cfg.Addr,
	}

	respChan := make(chan *pb.SubscribeResponse, 128)

	go func() {
		err = gnmi.SubscribeErr(ctx, client, &subOptions, respChan)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case response := <-respChan:
			switch resp := response.Response.(type) {
			case *pb.SubscribeResponse_Update:
				t := time.Unix(0, resp.Update.Timestamp).UTC()
				prefix := gnmi.StrPath(resp.Update.Prefix)
				var target string
				if t := resp.Update.Prefix.GetTarget(); t != "" {
					target = "(" + t + ") "
				}
				for _, update := range resp.Update.Update {
					fmt.Printf("[%s] %sUpdate %s = %s\n",
						t.Format(time.RFC3339Nano),
						target,
						path.Join(prefix, gnmi.StrPath(update.Path)),
						gnmi.StrUpdateVal(update),
					)
				}
			}
		}
	}

}
