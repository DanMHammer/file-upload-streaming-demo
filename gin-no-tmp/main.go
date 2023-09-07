package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/", func(c *gin.Context) {
		err := upload(c.Request)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "file uploaded successfully")
	})
	router.Run(":8080")
}

func WriteFileChunk(chunk *multipart.Part, file *os.File) error {
	buffer := make([]byte, 4096)
	bufbytes, err := chunk.Read(buffer)
	if err == io.EOF {
		return err
	}
	file.Write(buffer[:bufbytes])
	return err
}

func LargeFileStream(r *multipart.Reader) {
	var dir = os.Getenv("DIR")
	if dir == "" {
		dir = "/Users/svetlinmladenov/mnt"
	}

	p, err := r.NextPart()
	if err == io.EOF {
		return
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", dir, p.FileName()), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	for {
		err := WriteFileChunk(p, file)
		if err == io.EOF {
			return
		}
	}
}

func upload(r *http.Request) error {
	reader, err := r.MultipartReader()
	if err != nil {
		return err
	}

	LargeFileStream(reader)

	return nil
}
