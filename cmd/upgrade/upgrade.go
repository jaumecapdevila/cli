package upgrade

import (
	"github.com/dnote-io/cli/core"
	"github.com/dnote-io/cli/infra"
	"github.com/dnote-io/cli/upgrade"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var example = `
 dnote upgrade`

func NewCmd(ctx infra.DnoteCtx) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "upgrade",
		Short:   "Upgrades dnote",
		Example: example,
		RunE:    newRun(ctx),
	}

	return cmd
}

func newRun(ctx infra.DnoteCtx) core.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		if err := upgrade.Upgrade(ctx); err != nil {
			return errors.Wrap(err, "Failed to upgrade dnote")
		}

		return nil
	}
}
