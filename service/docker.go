package service

import (
	"context"
	"crimson/global"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type DockerApi struct {
}

func (a *DockerApi) CreateContainer() string {
	id := uuid.New().String()
	create, err := global.Docker.ContainerCreate(context.Background(), &container.Config{
		Image:        "code-server",
		Env:          []string{"PASSWORD=password123"}, // 设置密码的环境变量
		ExposedPorts: nat.PortSet{},
		User:         "root",
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"8080/tcp": []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
				},
			},
		},
	}, &network.NetworkingConfig{}, &ocispec.Platform{}, id)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return create.ID
}

func (d *DockerApi) StartContainer(containerId string) {
	err := global.Docker.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (d *DockerApi) StopContainer() {
	err := global.Docker.ContainerStop(context.Background(), "", container.StopOptions{})
	if err != nil {
		return
	}
}

func (d *DockerApi) PullConImages() {
	_, err := global.Docker.ImagePull(context.Background(), "codercom/code-server", types.ImagePullOptions{
		All:           false,
		RegistryAuth:  "",
		PrivilegeFunc: nil,
		Platform:      "",
	})
	if err != nil {
		return
	}
}
