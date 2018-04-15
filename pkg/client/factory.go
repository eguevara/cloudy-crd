package client

import (
	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"k8s.io/client-go/kubernetes"

	clientset "github.com/eguevara/cloudy-crd/pkg/generated/clientset/versioned"
)

type Factory interface {
	BindFlags(flags *pflag.FlagSet)
	Client() (clientset.Interface, error)
	KubeClient() (kubernetes.Interface, error)
	Namespace() string
}

type factory struct {
	flags      *pflag.FlagSet
	kubeconfig string
	baseName   string
	namespace  string
}

func NewFactory(baseName string) Factory {
	f := &factory{
		flags:    pflag.NewFlagSet("", pflag.ContinueOnError),
		baseName: baseName,
	}

	f.flags.StringVar(&f.kubeconfig, "kubeconfig", "", "Path to the kubeconfig file to use to talk to the Kubernetes apiserver. If unset, try the environment variable KUBECONFIG, as well as in-cluster configuration")
	f.flags.StringVarP(&f.namespace, "namespace", "n", f.namespace, "The namespace in which cloudy should operate")

	return f
}

func (f *factory) BindFlags(flags *pflag.FlagSet) {
	flags.AddFlagSet(f.flags)
}

func (f *factory) Client() (clientset.Interface, error) {
	clientConfig, err := Config(f.kubeconfig, f.baseName)
	if err != nil {
		return nil, err
	}

	cloudyClient, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return cloudyClient, nil
}

func (f *factory) KubeClient() (kubernetes.Interface, error) {
	clientConfig, err := Config(f.kubeconfig, f.baseName)
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return kubeClient, nil
}

func (f *factory) Namespace() string {
	return f.namespace
}
