package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload-file", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseMultipartForm(10 << 20)
		if err != nil {
			_, _ = fmt.Fprintf(writer, "Failed to parse multipart form:%s", err.Error())
			return
		}

		fileHeaders := request.MultipartForm.File["upload_file"][0]

		file, err := fileHeaders.Open()

		if err != nil {
			log.Fatalf("failed to open the file:%s", err.Error())
			return
		}

		//如果不分配[]byte的内存，只是创建了data []byte，那么read方法是没办法读取数据的哦
		//var data []byte
		data := make([]byte, 100)

		length, err := file.Read(data)

		if err != nil {
			log.Fatalf("failed to read data from the file:%s", err.Error())
		}

		log.Printf("The name of the file:%s,the size of the file:%d,the content of the file:%s,the read length:%d\n",
			fileHeaders.Filename,
			fileHeaders.Size,
			data,
			length,
		)

		_, err = fmt.Fprintf(writer, "the data of the file:%c\n", data)

		if err != nil {
			log.Fatalf("failed to return :%s", err.Error())
		}

	})

	err := http.ListenAndServe("0.0.0.0:9000", nil)

	if err != nil {
		fmt.Printf("Failed to start server:%s\n", err)
	}
}
