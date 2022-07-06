package main

import (
	"flag"
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

func main() {
	url := flag.String("s", "https://www.wildberries.ru/", "site url")
	flag.Parse()
	if strings.Contains(*url, "https://") || strings.Contains(*url, "http://") {
		if err := downloadSite(*url); err != nil {
			fmt.Println(err)
			return
		}
		if err := parseResources(*url); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("invalid url")
		return
	}

}

// Функция для скачивания статики сайта
func downloadSite(site string) error {
	// Отправляем GET запрос
	resp, err := http.Get(site)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Проверяем, что статус запроса успешный
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response, status code: %d\n", resp.StatusCode)
	}
	// Записываем полученное тело запроса в файл html формата
	fileName := strings.Split(site, "/")[2] + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// Функция для парсинга ресурсов сайта
func parseResources(site string) error {
	x, err := goquery.ParseUrl(site)
	if err != nil {
		return err
	}
	// Получаем по ссылкам изображения и описания внешнего вида сайта
	for _, url := range x.Find("").Attrs("href") {
		var str []string
		switch {
		case strings.Contains(url, ".png"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".jpg"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".css"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		}
	}
	return nil
}

// Функция для загрузки изображений и описаний внешнего вида сайта
func downloadResources(fileName string, url string) error {
	// Получаем данные по URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаём файл с переданным именем и пишем в него полученные данные
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
