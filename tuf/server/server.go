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
package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/runtimes-common/tuf/config"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/gcs_lib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kms_lib"
)

func UpdateSecrets(config config.TUFConfig, rootKeyFile string, targetKeyFile string, snapshotKeyFile string) error {
	errorStr := make([]string, 0)
	if rootKeyFile != "" {

		errorStr := append(errorStr, uploadSecret(rootKeyFile, config).Error())
	}
	if targetKeyFile != "" {
		encyptedRootFileContents, err := kms_lib.Encrypt(config, "")
		errorStr := append(errorStr, err.Error())
	}
	if snapshotKeyFile != "" {
		encyptedRootFilContents, err := kms_lib.Encrypt(config, "")
		errorStr := append(errorStr, err.Error())
	}
	return fmt.Errorf("Encountered following errors %s", strings.Join(errorStr, "\n"))
}

func uploadSecret(file string, config config.TUFConfig) error {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	encyptedFileContents, err := kms_lib.Encrypt(config, string(text))
	tmpFile, errWrite := ioutil.TempFile("", "key")
	defer os.Remove(tmpFile.Name())
	if err != nil {
		return err
	}
	ioutil.WriteFile(tmpFile.Name(), string(encyptedFileContents), os.ModePerm)
	gcs_lib.Upload(config.GCSProjectId, config.GCSBucketId, t)
}
