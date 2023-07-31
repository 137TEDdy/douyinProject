/*
   @Author Ted
   @Since 2023/7/27 9:25
*/

package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"path/filepath"
)

type Minio struct {
	MinioClient  *minio.Client
	endpoint     string
	port         string
	VideoBuckets string
	PicBuckets   string
}

func Test() {
	videopath := "D:/1/sss/"
	filename := "xxxx"
	fileFinalPath := filepath.Join(videopath, filename)
	fmt.Println(fileFinalPath)
}
