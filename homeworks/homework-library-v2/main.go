package main

import (
	"fmt"
)

type Library struct {
	Books []Book
}

type Book struct {
	Title  string
	Author string
	Pages  int
	isRead bool
}

func (l *Library) Add(b Book) error {
	err := b.Validate()
	if err != nil {
		return err
	}

	for _, book := range l.Books {
		if b.Title == book.Title {
			return fmt.Errorf("книга \"%s\" уже есть", book.Title)
		}
	}

	l.Books = append(l.Books, b)
	return nil
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

func (b *Book) setRead(toggleRead bool) {
	b.isRead = toggleRead
}

func (l *Library) MarkRead(title string) error {
	for i, bookName := range l.Books {
		if bookName.Title == title {
			l.Books[i].setRead(true)

			return nil
		}
	}

	return fmt.Errorf("название \"%s\"не найдено", title)
}

func (l *Library) Stats() (total, read int) {
	for _, b := range l.Books {
		if b.isRead {
			read++
		}
		total++
	}
	return total, read
}

func (l *Library) ByAuthor() map[string][]Book {
	booksByAuthor := map[string][]Book{}

	for _, b := range l.Books {
		booksByAuthor[b.Author] = append(booksByAuthor[b.Author], b)
	}

	return booksByAuthor
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
	book4 := Book{
		Title:  "Крутая книга",
		Author: "Донован",
		Pages:  30,
		isRead: false,
	}
	book5 := Book{
		Title:  "Язык прогрмариования Go",
		Author: "Донован",
		Pages:  10,
		isRead: false,
	}

	cntRead := 0
	cntBooks := 0

	lib := Library{}

	for i, b := range []Book{book1, book2, book3, book4, book5} {
		if err := lib.Add(b); err != nil {
			fmt.Println("ошибка:", i, err)
		}
	}

	if err := lib.MarkRead("Крутая книга"); err != nil {
		fmt.Println("ошибка:", err)
	}

	for _, b := range lib.Books {
		if b.isRead {
			cntRead++
		}
		cntBooks++

		fmt.Println(b)
	}

	total, read := lib.Stats()
	fmt.Printf("Прочитано: %d из %d\n", read, total)

	for author, books := range lib.ByAuthor() {
		fmt.Println()
		fmt.Println(author, ": ")
		for _, book := range books {
			fmt.Println("\t", book)
		}
	}
}
