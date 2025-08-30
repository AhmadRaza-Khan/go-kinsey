package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-dilve/config"
	"github.com/go-dilve/models"

	"github.com/gin-gonic/gin"
	"github.com/jlaffaye/ftp"
	"github.com/joho/godotenv"
)

var (
	apiKey = os.Getenv("API_KEY")
	apiURL = os.Getenv("API_URL")
	server    = os.Getenv("SERVER")
	username  = os.Getenv("USERNAME")
	password  = os.Getenv("PASSWORD")
	remoteDir = "/images/1024x1024 Item Images"
	localDir  = "./downloaded_files/images"
)

func Download_file(c *gin.Context) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request failed: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data, status code: " + http.StatusText(resp.StatusCode)})
		return
	}

	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", "attachment; filename=\"data.json\"")
	c.Header("Content-Length", resp.Header.Get("Content-Length"))

	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write response: " + err.Error()})
		return
	}
}

func SaveProducts(c *gin.Context) {

	file, err := os.Open("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var products []models.Product
	json.Unmarshal(byteValue, &products)

	for _, product := range products {
		config.DB.Create(&product)
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d products uploaded successfully", len(products))})
}

func DownloadDirectory(c *gin.Context) {
	ftpClient, err := ftp.Dial(server)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to connect to FTP server: %v", err)})
		return
	}
	defer ftpClient.Quit()

	if err := ftpClient.Login(username, password); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to login to FTP server: %v", err)})
		return
	}

	if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to create local directory: %v", err)})
		return
	}

	files, err := ftpClient.List(remoteDir)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to list files: %v", err)})
		return
	}

	for _, file := range files {
		if file.Type == ftp.EntryTypeFile {
			localFilePath := filepath.Join(localDir, file.Name)

			if _, err := os.Stat(localFilePath); err == nil {
				fmt.Println("Skipping (already exists):", localFilePath)
				continue
			}

			remoteFilePath := remoteDir + "/" + file.Name
			fmt.Println("Downloading:", remoteFilePath)

			if err := downloadFile(ftpClient, remoteFilePath, localFilePath); err != nil {
				fmt.Printf("Failed to download file %s: %v\n", remoteFilePath, err)
				continue
			}

			fmt.Println("Downloaded:", localFilePath)
		}
	}

	c.JSON(200, gin.H{"message": "Files downloaded successfully!"})
}

func downloadFile(ftpClient *ftp.ServerConn, remoteFilePath, localFilePath string) error {
	resp, err := ftpClient.Retr(remoteFilePath)
	if err != nil {
		return fmt.Errorf("failed to retrieve file %s: %v", remoteFilePath, err)
	}
	defer resp.Close()

	localFile, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create local file %s: %v", localFilePath, err)
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, resp); err != nil {
		return fmt.Errorf("failed to write file %s: %v", localFilePath, err)
	}

	return nil
}
