// Copyright 2022 Jetpack Technologies Inc and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.jetpack.io/devbox"
)

func InitCmd() *cobra.Command {
	command := &cobra.Command{
		Use:  "init [<dir>]",
		Args: cobra.MaximumNArgs(1),
		RunE: runInitCmd,
	}
	return command
}

func runInitCmd(cmd *cobra.Command, args []string) error {
	path := pathArg(args)

	_, err := devbox.InitConfig(path)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
