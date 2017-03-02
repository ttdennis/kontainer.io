package abstraction

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	engine "github.com/docker/docker/client"
)

// DCli is an interface to abstract dockers engine api client
type DCli interface {
	// Every function below invokes engine api client's function and returns its error property
	ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error
	ContainerKill(ctx context.Context, container, signal string) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error)
	ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
}

type dcliAbstract struct {
	cli engine.Client
}

func (d dcliAbstract) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	return d.cli.ContainerStart(ctx, container, options)
}

func (d dcliAbstract) ContainerKill(ctx context.Context, container, signal string) error {
	return d.cli.ContainerKill(ctx, container, signal)
}

func (d dcliAbstract) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (string, error) {
	idStruct, err := d.cli.ContainerExecCreate(ctx, container, config)
	return idStruct.ID, err
}

func (d dcliAbstract) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	return d.cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, containerName)
}

func (d dcliAbstract) ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error) {
	return d.cli.ImageInspectWithRaw(ctx, imageID)
}
