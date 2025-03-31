package main

import (
    "context"
	"log"
	"fmt"

	"github.com/zxjia2002/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/zxjia2002/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
    "github.com/cloudwego/kitex/client"
    consul "github.com/kitex-contrib/registry-consul"
    
)

func main() {
    
    r, err := consul.NewConsulResolver("127.0.0.1:8500")
    if err != nil {
        log.Fatal(err)
    }
    client, err := echoservice.NewClient("gomall", client.WithResolver(r))
    if err != nil {
        log.Fatal(err)
    }
    res, err := client.Echo(context.TODO(), &pbapi.Request{Message: "hello world! I'm a micro service func!"})
	if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%v", res)
}
