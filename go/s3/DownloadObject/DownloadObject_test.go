// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/google/uuid"
)

type Config struct {
	Bucket   string `json:"Bucket"`
	Filename string `json:"Filename"`
}

var configFileName = "config.json"

var globalConfig Config

func populateConfiguration(t *testing.T) error {
	content, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	text := string(content)

	err = json.Unmarshal([]byte(text), &globalConfig)
	if err != nil {
		return err
	}

	t.Log("Bucket:   " + globalConfig.Bucket)
	t.Log("Filename: " + globalConfig.Filename)

	return nil
}

func createBucket(sess *session.Session, bucket *string) error {
	svc := s3.New(sess)

	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}

	return nil
}

func putFile(sess *session.Session, bucket *string, filename *string) error {
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Unable to open file " + *filename)
		return err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket,
		Key:    filename,
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

func deleteBucket(sess *session.Session, bucket *string) error {
	svc := s3.New(sess)

	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: bucket,
	})

	err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter)
	if err != nil {
		return err
	}

	_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}

	return nil
}

func TestDownloadObject(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	t.Log("Starting unit test at " + nowString)

	err := populateConfiguration(t)
	if err != nil {
		t.Fatal(err)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	createdBucket := false

	if globalConfig.Bucket == "" {
		id := uuid.New()
		globalConfig.Bucket = "test-bucket-" + id.String()

		err := createBucket(sess, &globalConfig.Bucket)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("Created bucket " + globalConfig.Bucket)
	}

	if globalConfig.Filename == "" {
		globalConfig.Filename = "dummy.txt"

		err := putFile(sess, &globalConfig.Bucket, &globalConfig.Filename)
		if err != nil {
			t.Log("You'll have to delete bucket " + globalConfig.Bucket + " yourself")
			t.Fatal(err)
		}

		err = os.Remove(globalConfig.Filename)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = DownloadObject(sess, &globalConfig.Filename, &globalConfig.Bucket)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Downloaded " + globalConfig.Filename + " from bucket " + globalConfig.Bucket)

	if createdBucket {
		err := deleteBucket(sess, &globalConfig.Bucket)
		if err != nil {
			t.Log("You'll have to delete bucket " + globalConfig.Bucket + " yourself")
		}

		t.Log("Deleted bucket " + globalConfig.Bucket)
	}
}
