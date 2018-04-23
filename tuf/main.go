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
package main

import (
	"github.com/GoogleCloudPlatform/runtimes-common/ctc_lib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/config"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kms_lib"
	"github.com/spf13/cobra"
)

// Flags
var projectId = "my-encryption-prject"
var keyId = "testkey"
var keyRingId = "testkeyring"
var location = "global"
var text string

// Command
var RootCommand = &ctc_lib.ContainerToolCommand{
	ContainerToolCommandBase: &ctc_lib.ContainerToolCommandBase{
		Command: &cobra.Command{
			Use: "Demo kms",
		},
		Phase:           "test",
		DefaultTemplate: "{{.}}",
	},
	RunO: func(command *cobra.Command, args []string) (interface{}, error) {
		config := config.TUFConfig{
			ProjectId:   projectId,
			KeyRingId:   keyRingId,
			Location:    location,
			CryptoKeyId: keyId,
		}
		encrypt_resp, err := kms_lib.Encrypt(config, text)
		ctc_lib.Log.Info(encrypt_resp)
		ctc_lib.Log.Info(err)
		if err != nil {
			return nil, err
		}
		decrypt_resp, err1 := kms_lib.Decrypt(config, encrypt_resp.Ciphertext)
		return decrypt_resp, err1
	},
}

func main() {
	ctc_lib.Execute(RootCommand)
}

func init() {
	RootCommand.PersistentFlags().StringVarP(&text, "plain-text", "p", "this is secret", "Text to encrypt using key.")
}
