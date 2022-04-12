package cli

import (
	"github.com/bernmarx/download-manager/internal/domain/manager"
	"github.com/urfave/cli/v2"
)

func Process(c *cli.Context) error {
	link := c.Args().Get(0)
	dest := c.Args().Get(1)

	err := manager.Download(link, dest)

	return err
}
