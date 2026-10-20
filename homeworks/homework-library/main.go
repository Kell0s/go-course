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
	if b.Title == "" {
		return fmt.Errorf("название не найдено")
	}
	if b.Pages <= 0 {
		return fmt.Errorf("неверное кол-во страниц")
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
		Pages:  10,
		isRead: false,
	}
	book3 := Book{
		Title:  "",
		Author: "Донован",
		Pages:  10,
		isRead: false,
	}

	library := []Book{book1, book2, book3}
	cntRead := 0
	cntBooks := 0

	for _, book := range library {
		err := book.Validate()
		if err != nil {
			fmt.Println("ошибка: ", err)

			continue
		}

		cntBooks++
		if book.isRead {
			cntRead++
		}

		fmt.Println(book.String())
	}
	fmt.Printf("Прочитано: %d из %d\n", cntRead, cntBooks)
}
