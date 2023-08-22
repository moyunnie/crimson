package initalize

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

func InitDocker() *client.Client {
	//

	cli, err := client.NewClientWithOpts(client.WithHost("tcp://192.168.223.144:2375"),
		client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	ping, err := cli.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(ping)
	return cli
}
