package plugin

import (
	"github.com/crawlab-team/crawlab-core/config"
	"github.com/crawlab-team/crawlab-core/constants"
	"github.com/crawlab-team/crawlab-core/entity"
	"github.com/crawlab-team/crawlab-core/grpc/client"
	"github.com/crawlab-team/crawlab-core/interfaces"
	"os"
)

func NewGrpcClient() (c interfaces.GrpcClient, err error) {
	// options
	var opts []client.Option

	// grpc address
	if os.Getenv("CRAWLAB_PLUGIN_GRPC_ADDRESS") == "" {
		address, err := entity.NewAddressFromString(os.Getenv("CRAWLAB_PLUGIN_GRPC_ADDRESS"))
		if err != nil {
			return nil, err
		}
		opts = append(opts, client.WithAddress(address))
	}

	// plugin type
	opts = append(opts, client.WithSubscribeType(constants.GrpcSubscribeTypePlugin))

	// grpc client
	c, err = client.GetClient(config.DefaultConfigPath, opts...)
	if err != nil {
		return nil, err
	}

	// initialize
	if err := c.Init(); err != nil {
		return nil, err
	}

	return c, nil
}
