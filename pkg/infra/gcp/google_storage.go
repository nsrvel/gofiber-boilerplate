package gcp

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/exception"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ClientUploader struct {
	cl         *storage.Client
	projectId  string
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func InitialUploadGoogleStorage(conf *config.Config, uploadPath string) {

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", conf.CloudStorage.GoogleStorage.GoogleCredentialsFile)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: conf.CloudStorage.GoogleStorage.GoogleCloudStorageBucket,
		projectId:  conf.CloudStorage.GoogleStorage.ProjectID,
		uploadPath: uploadPath,
	}
}

func UploadFile(conf *config.Config, log *logrus.Logger, file *multipart.FileHeader, param string) (string, *exception.Error) {

	initialUploadPath := fmt.Sprintf("%s/%s/", conf.CloudStorage.GoogleStorage.AppName, conf.App.Env)
	InitialUploadGoogleStorage(conf, initialUploadPath)

	uuidNew, err := utils.GenerateUUID()
	if err != nil {
		return "", exception.NewError(500, err.Error(), err.Error())
	}
	extentionFile := filepath.Ext(file.Filename)

	fileName := fmt.Sprintf("%s%s", uuidNew, extentionFile)
	fullPath := fmt.Sprintf("%s%s", param, fileName)

	blobFile, err := file.Open()
	if err != nil {
		return "", exception.NewError(500, "Failed to create blobFile.", "Gagal membuat blobFile.")
	}

	err = Uploader.uploadFile(blobFile, fullPath)
	if err != nil {
		return "", exception.NewError(500, fmt.Sprintf("Error upload image file %s, err : "+err.Error(), file.Filename), fmt.Sprintf("Error upload image file %s, err : "+err.Error(), file.Filename))
	}

	finalPath := fmt.Sprintf("%s/%s/%s%s", conf.CloudStorage.GoogleStorage.GoogleCloudStorageURL, conf.CloudStorage.GoogleStorage.GoogleCloudStorageBucket, initialUploadPath, fullPath)
	log.Info(fmt.Sprintf("Upload '%s", finalPath))
	return finalPath, nil
}

func DeleteFile(conf *config.Config, log *logrus.Logger, fullPath string) error {

	initialUploadPath := fmt.Sprintf("%s/%s/", conf.CloudStorage.GoogleStorage.AppName, conf.App.Env)
	InitialUploadGoogleStorage(conf, initialUploadPath)

	str := fmt.Sprintf("%s/%s/", conf.CloudStorage.GoogleStorage.GoogleCloudStorageURL, conf.CloudStorage.GoogleStorage.GoogleCloudStorageBucket)
	if strings.Contains(fullPath, str) {
		fullPath = strings.Replace(fullPath, str, "", 1)
	}
	errorDeleteFile := Uploader.deleteFile(fullPath)
	if errorDeleteFile != nil {
		return errorDeleteFile
	}
	log.Info(fmt.Sprintf("Delete '%s'", fullPath))
	return nil
}

func (c *ClientUploader) uploadFile(file multipart.File, object string) error {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (c *ClientUploader) deleteFile(object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := c.cl.Bucket(c.bucketName).Object(object)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("object.Attrs: %v", err)
	}

	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}
	return nil
}
