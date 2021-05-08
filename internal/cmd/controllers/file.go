package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
}

func AcquireFileName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"file_name": 5})
}

func GetFile(c *gin.Context) {
}

func GetFileMeta(c *gin.Context) {
}

func DeleteFile(c *gin.Context) {
}
