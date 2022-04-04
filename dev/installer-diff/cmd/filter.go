// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// filterCmd represents the inject command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// cfg, ns, err := getKubeconfig()
		// if err != nil {
		// 	log.WithError(err).Fatal("cannot get Kubernetes client config")
		// }
		// err = dart.Remove(cfg, ns, args[0])
		// if err != nil {
		// 	log.WithError(err).Fatal("cannot remove toxiproxy")
		// }
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
