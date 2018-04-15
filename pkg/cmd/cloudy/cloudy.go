package cloudy

import (
	"github.com/spf13/cobra"
	"github.com/eguevara/cloudy-crd/pkg/client"
	"github.com/eguevara/cloudy-crd/pkg/cmd/cli/deployment"
	"flag"
)

func NewCommand(name string) *cobra.Command {
	c := &cobra.Command{
		Use: name,
		Short: "Manage cloud commands.",
		Long: "Cloud tool is used to manage cloud resources like deploying releases.",
	}

	f := client.NewFactory(name)
	f.BindFlags(c.PersistentFlags())

	c.AddCommand(
		deployment.NewCommand(f),
	)

	c.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	flag.CommandLine.Parse([]string{})

	return c
}