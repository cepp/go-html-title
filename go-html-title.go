package gohtmltitle

import (
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

//  <- chan - canal somente-leitura
// Titulo extrai o título das urls enviadas via parâmetro
func Titulo(urls ...string) <-chan string {
	channel := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)

			read, _ := regexp.Compile("<title>(.*?)<\\/title>")
			channel <- read.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return channel
}
