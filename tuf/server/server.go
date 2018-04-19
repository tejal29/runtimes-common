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

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"github.com/GoogleCloudPlatform/runtimes-common/tuf/config"
	cloudkms "google.golang.org/api/cloudkms/v1"
)

func GetKmsService() {
	ctx := context.Background()
	config := config.TUFConfig{
		ProjectId: "my-encryption-prject",
		Location:  "global",
	}
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		fmt.Println(err)
	}
	kmsService, err := cloudkms.New(client)
	if err != nil {
		fmt.Println(err)
	}

	// The resource name of the key rings.
	parentName := fmt.Sprintf("projects/%s/locations/%s", config.ProjectId, config.Location)

	// Make the RPC call.
	response, err := kmsService.Projects.Locations.KeyRings.List(parentName).Do()
	if err != nil {
		fmt.Println("Failed to list key rings: %v", err)
	}

	// Print the returned key rings.
	for _, keyRing := range response.KeyRings {
		fmt.Printf("KeyRing: %q\n", keyRing.Name)
	}
}
