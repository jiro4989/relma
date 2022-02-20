package downloader

import (
	"net/http"
)

type DownloaderInterface interface {
	Download(string) (*http.Response, error)
}

type Downloader struct {
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) Download(url string) (*http.Response, error) {
	return http.Get(url)
}
