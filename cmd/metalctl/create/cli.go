package create

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create resources",
		Long:  "Commands that create resources in Equinix Metal",
	}

	return command
}
