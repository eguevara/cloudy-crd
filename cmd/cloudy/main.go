package main

import (
	"github.com/golang/glog"
	"path/filepath"
	"os"
	"github.com/eguevara/cloudy-crd/pkg/cmd/cloudy"
)

func main() {
	defer glog.Flush()

	baseName := filepath.Base(os.Args[0])

	_ = cloudy.NewCommand(baseName).Execute()

}