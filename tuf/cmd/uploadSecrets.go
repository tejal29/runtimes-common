/*
Copyright 2018 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/GoogleCloudPlatform/runtimes-common/ctc_lib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/server"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/util"

	"github.com/spf13/cobra"
)

// Flags
var rootKey = "~/root.json"
var targetKey = "~/target.json"
var snapshotKey = "~/snapshot.json"

// Command To upload Secrets to GCS store and Renegerate all the MetaData.
var UploadSecretsCommand = &ctc_lib.ContainerToolCommand{
	ContainerToolCommandBase: &ctc_lib.ContainerToolCommandBase{
		Command: &cobra.Command{
			Use: "Prototype GCS Update Keys to Google Cloud Storage.",
		},
		Phase:           "test",
		DefaultTemplate: "{{.}}",
	},
	RunO: func(command *cobra.Command, args []string) (interface{}, error) {
		tufConfig, err := util.ReadConfig(tufConfigFilename)
		if err != nil {
			return nil, err
		}
		server.UpdateSecrets(tufConfig, rootKey, targetKey, snapshotKey)
		return nil, nil
	},
}

func init() {
	RootCommand.PersistentFlags().StringVar(&rootKey, "root-key", "", "GCloud key.json for Root role")
	RootCommand.PersistentFlags().StringVar(&targetKey, "target-key", "", "GCloud key.json for Snapshot role")
	RootCommand.PersistentFlags().StringVar(&snapshotKey, "snapshot-key", "", "GCloud key.json for Target role")
}
