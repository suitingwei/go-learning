package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

const basePath = "/Users/didi/Downloads/new-world"
const targetPath = "/tmp/summary"

//每10个目录，合并一次
const fileStep = 10

func main() {
	rootDirInfo, err := ioutil.ReadDir(basePath)

	if err != nil {
		fmt.Println("Failed to open file", err)
		os.Exit(-1)
	}

	createDir(targetPath, true)

	fileIndex := 1

	for _, subDir := range rootDirInfo {
		if !subDir.IsDir() {
			fmt.Printf("SubDir=%s is not dir,skip it\n", subDir.Name())
			continue
		}

		targetDirName := getTargetDirName(fileIndex)
		fmt.Printf("Processing sub dir=%s,index=%d,targetDirName=%s\n", subDir.Name(), fileIndex, targetDirName)

		//目标目录(已经分好组)
		targetDirPath := path.Join(targetPath, targetDirName)
		createDir(targetDirPath, false)

		mvFiles(subDir, targetDirPath)

		fileIndex++
	}
	fmt.Println()
}

//按照10个目录，生成一个分组
//1~10 --> "1~10"
//11-20 --> "11~20"
//41~50 ---> 41~50
func getTargetDirName(fileIndex int) string {
	start := fileIndex/10*10 + 1
	end := fileIndex/10*10 + 10

	//fmt.Printf("----Start=%d,end=%d\n", start, end)

	return fmt.Sprintf("%d~%d", start, end)
}

func mvFiles(subDir os.FileInfo, targetDirPath string) {
	subDirPath := path.Join(basePath, subDir.Name())
	subDirInfo, err := ioutil.ReadDir(subDirPath)
	if err != nil {
		fmt.Println("Failed to open sub dir,", err)
		os.Exit(-1)
	}
	for _, file := range subDirInfo {
		if file.IsDir() {
			fmt.Printf("\tDir=%s is not file,skip it\n", file.Name())
			continue
		}
		originalFileName := path.Join(subDirPath, file.Name())
		targetFileName := path.Join(targetDirPath, subDir.Name()+"_"+file.Name())
		fmt.Printf("From=%s,to=%s\n", originalFileName, targetFileName)

		_, err = copyFile(originalFileName, targetFileName)

		if err != nil {
			fmt.Println("Failed to copy file", err)
			os.Exit(-1)
		}
	}
}

func createDir(dirPath string, forceDelete bool) {
	if forceDelete && Exists(dirPath) {
		err := os.RemoveAll(dirPath)
		if err != nil {
			fmt.Println("Failed to delete path", err)
		}
	}
	err := os.Mkdir(dirPath, os.ModeDir|os.ModePerm)
	if err != nil {
		if os.IsExist(err) && !forceDelete {
			return
		}
		fmt.Printf("Failed to create path=%s,err=%s\n", dirPath, err.Error())
		os.Exit(-1)
	} else {
		//fmt.Printf("Create dir=%s successfully\n", dirPath)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer func() {
		err = src.Close()
		if err != nil {
			return
		}
	}()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer func() {
		err = dst.Close()
		if err != nil {
			return
		}
	}()
	return io.Copy(dst, src)
}
