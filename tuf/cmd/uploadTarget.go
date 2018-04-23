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
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/config"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kms_lib"
	"github.com/spf13/cobra"
)

// Command
var UploadTargetCommand = &ctc_lib.ContainerToolCommand{
	ContainerToolCommandBase: &ctc_lib.ContainerToolCommandBase{
		Command: &cobra.Command{
			Use: "Prototype GCS Update Targets to Google Cloud Storage.",
		},
		Phase:           "test",
		DefaultTemplate: "{{.}}",
	},
	RunO: func(command *cobra.Command, args []string) (interface{}, error) {
		config := config.TUFConfig{
			GCSProjectId: GCSProjectId,
			GCSBucketId:  GCSBucketId,
			KMSProjectId: KMSProjectId,
			KeyRingId:    KeyRingId,
			KMSLocation:  KMSLocation,
			CryptoKeyId:  KeyId,
		}
		encrypt_resp, err := kms_lib.Encrypt(config, "")
		if err != nil {
			return nil, err
		}
		decrypt_resp, dispErr := kms_lib.Decrypt(config, encrypt_resp.Ciphertext)
		return decrypt_resp, dispErr
	},
}
