// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	// "gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	// "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	// // "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

// filterCmd represents the inject command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		filter(cmd, args)
	},
}

func filter(_ *cobra.Command, _ []string) {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read from stdin: %s", err)
	}

	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(input, nil, nil)
	if err != nil {
		log.Fatalf("failed to deserialize: %s", err)
	}

	recurse(obj)
}

func recurse(obj runtime.Object) {
	switch v := obj.(type) {
	case *corev1.List:
		fmt.Printf("object is a list\n")
		for _, item := range v.Items {
			fmt.Printf("one child: %v\n", item.Object)
		}
	case *appsv1.Deployment:
		fmt.Printf("object is a deployment\n")
		// default:
		// 	fmt.Printf("object is something else\n")
	}
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
