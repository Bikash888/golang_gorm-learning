package utils

import (
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

//Uploads Files in Google Console
func HandleImageInUploadToGCP(c *gin.Context) {
	bucket := "test-loc-images"
	var err error
	ctx := appengine.NewContext(c.Request)
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("serviceKey.json"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadImageFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	defer f.Close()

	serviceWorker := storageClient.Bucket(bucket).Object(uploadImageFile.Filename).NewWriter(ctx)
	if _, err = io.Copy(serviceWorker, f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	if err := serviceWorker.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	u, err := url.Parse("/" + bucket + "/" + serviceWorker.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}
