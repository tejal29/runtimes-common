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
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kmsLib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/util"
	"github.com/spf13/cobra"
)

// Flags
var tufConfigFilename string

// Command
var RootCommand = &ctc_lib.ContainerToolCommand{
	ContainerToolCommandBase: &ctc_lib.ContainerToolCommandBase{
		Command: &cobra.Command{
			Use: "Prototype GCS TuF",
		},
		Phase:           "test",
		DefaultTemplate: "{{.}}",
	},
	RunO: func(command *cobra.Command, args []string) (interface{}, error) {
		tufConfig, err := util.ReadConfig(tufConfigFilename)
		encryptResp, err := kmsLib.Encrypt(tufConfig, "")
		if err != nil {
			return nil, err
		}
		decryptResp, dispErr := kmsLib.Decrypt(tufConfig, encryptResp.Ciphertext)
		return decryptResp, dispErr
	},
}

func init() {
	RootCommand.PersistentFlags().StringVarP(&tufConfigFilename, "config", "c", "tuf.yaml", "File name for Tool config")
	RootCommand.AddCommand(UploadSecretsCommand)
	RootCommand.AddCommand(UploadTargetCommand)
	RootCommand.AddCommand(GenerateKeysCommand)
}
