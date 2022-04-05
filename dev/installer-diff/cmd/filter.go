// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"
	"io/fs"
	// "io/ioutil"
	// "log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	// "gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	// "k8s.io/apimachinery/pkg/api/meta"
	// corev1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// filterCmd represents the inject command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		filter(args[0])
	},
}

func filter(dir string) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		fmt.Printf("reading %s\n", path)
		yaml, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read file: %s\n", path)
			return nil
		}

		obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(yaml, nil, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to deserialize: %s\n", err)
			return nil
		}

		handle(obj)
		return nil
	})
}

func handle(obj runtime.Object) {
	switch v := obj.(type) {
	case *appsv1.Deployment:
		fmt.Printf("object is a deployment: %s\n", v.Name)
	default:
		fmt.Printf("object is something else\n")
	}
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
