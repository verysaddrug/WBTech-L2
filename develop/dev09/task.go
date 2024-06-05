package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wget <URL>")
		os.Exit(1)
	}

	startURL := os.Args[1]

	// Скачивание начальной страницы
	err := downloadPage(startURL, ".")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func downloadPage(pageURL string, baseDir string) error {
	// Отправка HTTP-запроса
	resp, err := http.Get(pageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверка успешного ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download page: %s", resp.Status)
	}

	// Создание файла для сохранения
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return err
	}
	filePath := filepath.Join(baseDir, parsedURL.Host, parsedURL.Path)
	if filePath[len(filePath)-1] == '/' {
		filePath += "index.html"
	}

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Запись содержимого в файл
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = outFile.Write(body)
	if err != nil {
		return err
	}

	// Поиск и скачивание связанных страниц
	relatedLinks, err := extractLinks(strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	for _, link := range relatedLinks {
		absoluteLink := toAbsoluteURL(link, pageURL)
		err := downloadPage(absoluteLink, baseDir)
		if err != nil {
			fmt.Println("Warning:", err)
		}
	}

	return nil
}

func extractLinks(body io.Reader) ([]string, error) {
	var links []string
	tokenizer := html.NewTokenizer(body)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				return links, nil
			}
			return nil, tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func toAbsoluteURL(href string, base string) string {
	parsedHref, err := url.Parse(href)
	if err != nil {
		return href
	}
	if parsedHref.IsAbs() {
		return href
	}
	parsedBase, err := url.Parse(base)
	if err != nil {
		return href
	}
	return parsedBase.ResolveReference(parsedHref).String()
}
