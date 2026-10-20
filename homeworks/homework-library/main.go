package main

import "fmt"

type book struct {
	Title  string
	Author string
	Pages  int
	isRead bool
}

func String(b book) string {
	if b.isRead {
		return fmt.Sprintf("«%s» - %s, %d стр ✓",
			b.Title, b.Author, b.Pages)
	}

	return fmt.Sprintf("«%s» - %s, %d стр",
		b.Title, b.Author, b.Pages)
}

func Validate(b book) error {
	errTitle := "Ошибка: нет названия"
	errPage := "Ошибка: неверное кол-во страниц"

	if b.Title == "" {
		return fmt.Errorf("%s", errTitle)
	}
	if b.Pages <= 0 {
		return fmt.Errorf("%s", errPage)
	}

	return nil
}

func main() {
	book1 := book{
		Title:  "",
		Author: "Мартин",
		Pages:  400,
		isRead: true,
	}
	book2 := book{
		Title:  "Язык прогрмариования Go",
		Author: "Донован",
		Pages:  -10,
		isRead: false,
	}

	library := []book{book1, book2}

	for _, book := range library {
		if Validate(book) == nil {
			fmt.Println(String(book))
		} else {
			fmt.Println(Validate(book))
		}
	}
}
