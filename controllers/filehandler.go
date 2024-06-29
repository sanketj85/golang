package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadFile handles file upload
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a random filename
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": filename})
}

// getUploadedFiles retrieves list of uploaded files
func GetUploadedFiles(c *gin.Context) {
	files := []string{}

	err := filepath.Walk("./uploads", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve files"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func DownloadFile(c *gin.Context) {
    filename := c.Param("filename")
    fileLocation := "./uploads/" + filename

    // Check if file exists
    _, err := os.Stat(fileLocation)
    if os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
        return
    }

    // Serve file as attachment
    c.Header("Content-Disposition", "attachment; filename="+filename)
    c.File(fileLocation)
}
