package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
	isRead bool
}

func (b Book) String() string {
	if b.isRead {
		return fmt.Sprintf("«%s» - %s, %d стр ✓",
			b.Title, b.Author, b.Pages)
	}

	return fmt.Sprintf("«%s» - %s, %d стр",
		b.Title, b.Author, b.Pages)
}

func (b Book) Validate() error {
	errTitle := "нет названия"
	errPage := "неверное кол-во страниц"

	if b.Title == "" {
		return fmt.Errorf("ошибка: %s", errTitle)
	}
	if b.Pages <= 0 {
		return fmt.Errorf("ошибка: %s", errPage)
	}

	return nil
}

func main() {
	book1 := Book{
		Title:  "Чистый код",
		Author: "Мартин",
		Pages:  400,
		isRead: true,
	}
	book2 := Book{
		Title:  "Язык прогрмариования Go",
		Author: "Донован",
		Pages:  -10,
		isRead: false,
	}

	library := []Book{book1, book2}

	for _, book := range library {
		err := book.Validate()
		if err != nil {
			fmt.Println(book.Validate())
			continue
		}

		fmt.Println(book.String())
	}
}
