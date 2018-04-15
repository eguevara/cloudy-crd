package deployment

import (
	"fmt"
	"github.com/eguevara/cloudy-crd/pkg/client"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewGetCommand(f client.Factory, use string) *cobra.Command {
	var listOptions metav1.ListOptions

	c := &cobra.Command{
		Use:   use,
		Short: "Get deployments",
		Run: func(c *cobra.Command, args []string) {

			cloudyClient, _ := f.Client()

			deployments,err  := cloudyClient.CloudyV1().Deployments("default").List(listOptions)
			if err != nil {
				fmt.Println(err)
			}

			for _, db := range deployments.Items {
				fmt.Printf("deployment %s on %s\n", db.Name, db.Spec.Environment)
			}

		},
	}

	f.BindFlags(c.Flags())

	return c
}
