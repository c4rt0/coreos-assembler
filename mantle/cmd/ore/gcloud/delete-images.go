// Copyright 2017 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcloud

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/api/googleapi"

	"github.com/coreos/coreos-assembler/mantle/platform/api/gcloud"
)

var (
	cmdDeleteImage = &cobra.Command{
		Use:   "delete-images <name>...",
		Short: "Delete GCP images",
		Run:   runDeleteImage,
	}
	allowMissing bool
)

func init() {
	GCloud.AddCommand(cmdDeleteImage)
	cmdDeleteImage.Flags().BoolVar(&allowMissing, "allow-missing", false, "Do not error out on the resource not existing")
}

func runDeleteImage(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Specify image name(s).\n")
		os.Exit(2)
	}

	exit := 0
	pendings := map[string]*gcloud.Pending{}
	for _, name := range args {
		pending, err := api.DeleteImage(name)
		if err != nil {
			if gErr, ok := err.(*googleapi.Error); ok {
				// Skip on NotFound error only if allowMissing flag is set to True
				if gErr.Code == 404 && allowMissing {
					plog.Infof("%v\n", err)
					continue
				}
			}
			fmt.Fprintf(os.Stderr, "Deleting %q failed: %v\n", name, err)
			exit = 1
			continue
		}
		pendings[name] = pending
	}
	for name, pending := range pendings {
		if err := pending.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "Deleting %q failed: %v\n", name, err)
			exit = 1
		}
	}
	os.Exit(exit)
}
