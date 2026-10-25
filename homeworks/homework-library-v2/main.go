package main

import "fmt"

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
			return fmt.Errorf("книга %s уже есть", book)
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

	book2.setRead(true)

	lib := Library{
		Books: []Book{book1, book2, book3},
	}

	err := lib.Add(book4)
	if err != nil {
		fmt.Println("ошибка:", err)
	}

	err = lib.Add(book5)
	if err != nil {
		fmt.Println("ошибка:", err)
	}

	cntRead := 0
	cntBooks := 0

	for _, book := range lib.Books {
		err := book.Validate()
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}

		cntBooks++
		if book.isRead {
			cntRead++
		}

		fmt.Println(book)
	}
	fmt.Printf("Прочитано: %d из %d\n", cntRead, cntBooks)
}
