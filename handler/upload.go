package handler

import (
	"flash-sale-backend/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	v, _ := c.Get("config")
	config := v.(*utils.MyConfig)
	// Source
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	filename := filepath.Base(file.Filename)
	filepath := filepath.Join(config.ServerConfig.UploadDir, filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"url":      config.ServerConfig.ImagePrefix + filename,
	})
}
