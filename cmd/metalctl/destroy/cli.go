package destroy

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "destroy",
		Short: "Destroy resources",
		Long:  "Commands that destroy resources in Equinix Metal",
	}

	return command
}
