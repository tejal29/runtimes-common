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
package gcsLib

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"

	"golang.org/x/net/context"
)

func Upload(projectId string, bucket string, name string, public bool, r io.Reader) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	bh := client.Bucket(bucket)
	// Next check if the bucket exists
	if _, err = bh.Attrs(ctx); err != nil {
		// TODO: Create a new bucket.
		return nil, nil, errors.New("Please create the bucket first e.g. with `gsutil mb`")
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}

func Download(projectId string, bucketId string, objectName string, destPath string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bh := client.Bucket(bucketId)

	rc, err := bh.Object(objectName).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(destPath, slurp, os.ModePerm) // Create a file with 777 mode.
}
