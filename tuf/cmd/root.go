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
	"github.com/spf13/viper"
)

// Flags
var GCSProjectId = "my-gcs-prject"
var KeyId = "testkey"
var KeyRingId = "testkeyring"
var KMSLocation = "global"
var KMSProjectId = "my-encryption-prject"
var GCSBucketId = "test-bucket"

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

func init() {
	RootCommand.PersistentFlags().StringVar(&GCSBucketId, "gcs_bucket_id", "<GCS_BUCKET_ID>", "Google Cloud Storage Bucket id.")
	viper.BindPFlag("gcsBucketId", RootCommand.PersistentFlags().Lookup("gcs_bucket_id"))
	RootCommand.PersistentFlags().StringVar(&GCSProjectId, "gcs_project_id", "<GCS_PROJECT_ID>", "Google Cloud Storage Project Id.")
	viper.BindPFlag("gcs_project_id", RootCommand.PersistentFlags().Lookup("gcs_project_id"))
	RootCommand.PersistentFlags().StringVar(&KMSProjectId, "kms_project_id", "<KMS_PROJECT_ID>", "Google KMS Storage Project Id")
	viper.BindPFlag("kms_project_id", RootCommand.PersistentFlags().Lookup("kms_project_id"))
	RootCommand.PersistentFlags().StringVar(&KeyRingId, "keyringid", "<KMS_KEYRING_ID>", "Google KMS ring id.")
	viper.BindPFlag("keyringid", RootCommand.PersistentFlags().Lookup("keyringid"))
	RootCommand.PersistentFlags().StringVar(&KeyId, "cryptokey", "<KMS_CRYPTO_KEY>", "Google KMS key id")
	viper.BindPFlag("cryptokey", RootCommand.PersistentFlags().Lookup("cryptokey"))

	// Add Subcommands
	RootCommand.AddCommand(UploadSecretsCommand)
	RootCommand.AddCommand(UploadTargetCommand)
	RootCommand.AddCommand(GenerateKeysCommand)
}
