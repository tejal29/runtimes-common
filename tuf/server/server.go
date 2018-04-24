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

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/config"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/gcsLib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/kmsLib"
	"github.com/GoogleCloudPlatform/runtimes-common/tuf/metadata/v1"
)

type TUFMetadata struct {
	RootFile v1.Metadata
	Target   v1.Metadata
	Snapshot v1.Metadata
}

type KeyPair struct {
	Public  string
	Private string
}

func UpdateSecrets(tufConfig config.TUFConfig, rootKeyFile string, targetKeyFile string, snapshotKeyFile string) error {
	errorStr := make([]string, 0)
	oldRootKeyFile := ""
	if rootKeyFile != "" {
		// in case of Root Key being updated, we need to first download the old root key and sign the root.json with
		// old and new key.

		tmpFile, errWrite := ioutil.TempFile("", "root.key.old")
		defer os.Remove(tmpFile.Name())
		if errWrite != nil {
			return errWrite
		}
		err := gcsLib.Download(tufConfig.GCSProjectId, tufConfig.GCSBucketId, config.RootSecretFileName, tmpFile.Name())
		if err != nil && err != storage.ErrObjectNotExist {
			// The old root file exists but there was an error reading it. This is fatal hence return error
			return err
		}
		oldRootKeyFile = tmpFile.Name()
		errorStr = append(errorStr, uploadSecret(rootKeyFile, tufConfig, config.RootSecretFileName).Error())
	}
	if targetKeyFile != "" {
		errorStr = append(errorStr, uploadSecret(targetKeyFile, tufConfig, config.TargetSecretFileName).Error())
	}
	if snapshotKeyFile != "" {
		errorStr = append(errorStr, uploadSecret(snapshotKeyFile, tufConfig, config.SnapshotSecretFileName).Error())
	}
	if len(errorStr) > 1 {
		// Exit if there were errors uploading secrets.
		return fmt.Errorf("Encountered following errors %s", strings.Join(errorStr, "\n"))
	}

	// Generate all the Metadata.
	tufMetadata := GenerateMetadata(rootKeyFile, oldRootKeyFile, targetKeyFile, snapshotKeyFile)

	// Write Consistent Snapshots
	WriteConsistentSnapshot(tufMetadata)
	return nil
}

func uploadSecret(file string, tufConfig config.TUFConfig, name string) error {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	encyptedResponse, err := kmsLib.Encrypt(tufConfig, string(text))
	tmpFile, errWrite := ioutil.TempFile("", "key")
	defer os.Remove(tmpFile.Name())
	if errWrite != nil {
		return err
	}
	ioutil.WriteFile(tmpFile.Name(), []byte(encyptedResponse.Ciphertext), os.ModePerm)
	tmpFile.Close()
	_, _, err = gcsLib.Upload(tufConfig.GCSProjectId, tufConfig.GCSBucketId, name, false, tmpFile)
	return err
}

func GenerateMetadata(rootKeyFile string, oldRootKeyFile string, targetKeyFile string, snapshotKeyFile string) *TUFMetadata {
	// Populate Root.json Signed part

	// Sign with new key and populate the unsigned part

	if oldRootKeyFile != "" {
		// Sign with old key and add this to unsigned part
	}

	// Write root.json and <n>/root.json
	// Push it.

	// Write <n>/target.json
	//Push it

	// Write <n>/Snapshot.json
	//Push it
	return &TUFMetadata{}
}

func WriteConsistentSnapshot(tufMetadata *TUFMetadata) {

}
