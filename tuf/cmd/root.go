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
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kmsLib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags
var GCSProjectId = "tuf-storage-project"
var KeyId = "testkey"
var KeyRingId = "testkeyring"
var KMSLocation = "global"
var KMSProjectId = "my-encryption-prject"
var GCSBucketId = "tuf-store"

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
		encrypt_resp, err := kmsLib.Encrypt(config, "")
		if err != nil {
			return nil, err
		}
		decrypt_resp, dispErr := kmsLib.Decrypt(config, encrypt_resp.Ciphertext)
		return decrypt_resp, dispErr
	},
}

func init() {
	RootCommand.PersistentFlags().StringVar(&GCSBucketId, "gcs-bucket-id", "<GCS_BUCKET_ID>", "Google Cloud Storage Bucket id.")
	viper.BindPFlag("gcsBucketId", RootCommand.PersistentFlags().Lookup("gcs-bucket-id"))
	RootCommand.PersistentFlags().StringVar(&GCSProjectId, "gcs-project-id", "<GCS_PROJECT_ID>", "Google Cloud Storage Project Id.")
	viper.BindPFlag("gcsProjectId", RootCommand.PersistentFlags().Lookup("gcs-project-id"))
	RootCommand.PersistentFlags().StringVar(&KMSProjectId, "kms-project-id", "<KMS_PROJECT_ID>", "Google KMS Storage Project Id")
	viper.BindPFlag("kmsProjectId", RootCommand.PersistentFlags().Lookup("kms-project-id"))
	RootCommand.PersistentFlags().StringVar(&KeyRingId, "keyring-id", "<KMS_KEYRING_ID>", "Google KMS ring id.")
	viper.BindPFlag("keyringId", RootCommand.PersistentFlags().Lookup("keyring-id"))
	RootCommand.PersistentFlags().StringVar(&KeyId, "crypto-key", "<KMS_CRYPTO_KEY>", "Google KMS key id")
	viper.BindPFlag("cryptoKey", RootCommand.PersistentFlags().Lookup("crypto-key"))

	// Add Subcommands
	RootCommand.AddCommand(UploadSecretsCommand)
	RootCommand.AddCommand(UploadTargetCommand)
	RootCommand.AddCommand(GenerateKeysCommand)
}
