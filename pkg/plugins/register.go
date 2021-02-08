package plugins

import (
	"github.com/morningfish/kube-scheduler-framework/pkg/plugins/physicalnode"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(physicalnode.Name, physicalnode.New),
	)
}
