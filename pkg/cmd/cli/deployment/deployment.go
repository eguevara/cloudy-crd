package deployment

import (
	"github.com/eguevara/cloudy-crd/pkg/client"
	"github.com/spf13/cobra"
)

func NewCommand(f client.Factory) *cobra.Command {
	c := &cobra.Command{
		Use:   "deployment",
		Short: "Work with deployments",
		Long:  "Work with deployments",
	}


	c.AddCommand(
		NewGetCommand(f, "get"),
	)
	return c

}