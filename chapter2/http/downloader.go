package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type GithubUserInfo struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	Avatar string `json:"avatar_url"`
}

type GithubUserList struct {
	Users            []GithubUserInfo `json:"items"`
	TotalCount       int              `json:"total_count"`
	InCompleteResult bool             `json:"incomplete_results"`
}

type Downloader struct {
}

func New() *Downloader {
	return &Downloader{}
}

func (downloader Downloader) download1(users *GithubUserList) {

	//使用协程并发下载哦
	wg := sync.WaitGroup{}
	wg.Add(users.TotalCount)

	for _, user := range users.Users {
		//go func(user *GithubUserInfo) {
		//	downloadUrl(user, &wg)
		//}(&user)
		go downloadUrl(&user, &wg)
	}
	wg.Wait()
}

func downloadUrl(user *GithubUserInfo, wg *sync.WaitGroup) {
	fmt.Printf("\tuser=%T\n", user)
	defer wg.Done()
	imagePath, err := filepath.Abs("./")
	if err != nil {
		log.Fatalf("Failed to obtain the image path:%s\n", err.Error())
	}

	imagePath = filepath.Join(imagePath, "githubUserAvatars")

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		_ = os.Mkdir(imagePath, os.ModePerm)
	}

	file, err := os.Create(filepath.Join(imagePath, user.Login+".jpg"))

	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Get(user.Avatar)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User:%s's avatar:[%s] has been downloaded successfully!\n", user.Login, user.Avatar)
}
