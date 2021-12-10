package remote

import (
	"flag"

	ec "github.hpe.com/hpe/hpc-rabsw-nnf-ec/pkg/ec"
	nnf "github.hpe.com/hpe/hpc-rabsw-nnf-ec/pkg/manager-nnf"
	server "github.hpe.com/hpe/hpc-rabsw-nnf-ec/pkg/manager-server"
)

type Options struct{}

func BindFlags(fs *flag.FlagSet) *Options {
	return &Options{}
}

func RunController() error {

	opts := BindFlags(flag.CommandLine)

	flag.Parse()

	c := &ec.Controller{
		Name: "Near Node Flash Server",
		Routers: []ec.Router{
			nnf.NewDefaultApiRouter(
				nnf.NewDefaultApiService(
					NewDefaultServerStorageService(opts),
				), nil),
		},
	}

	err := c.Init(&ec.Options{
		Http:    true,
		Port:    server.RemoteStorageServicePort,
		Log:     true,
		Verbose: true,
	})

	if err != nil {
		return err
	}

	c.Run()

	return nil
}
