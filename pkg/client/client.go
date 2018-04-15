package client

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

)

func Config(kubeconfig, baseName string) (*rest.Config, error) {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	loader.ExplicitPath = kubeconfig
	clientConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", loader.Load)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	clientConfig.UserAgent = buildUserAgent(
		baseName,
		runtime.GOOS,
		runtime.GOARCH,
	)

	return clientConfig, nil
}

func buildUserAgent(command, os, arch string) string {
	return fmt.Sprintf(
		"%s/%s (%s/%s) %s", command, os, arch )
}
