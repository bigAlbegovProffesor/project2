package domain
import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int //ID читателя, который взял книгу
}

// IssueBook выдает книгу читателю. Теперь возвращает ошибку.
func (b *Book) IssueBook(reader *Reader) error {
	if b.IsIssued {
		//Теперь возвращаем ошибку, а не печатаем в консоль
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу.", reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	//fmt.Println больше не нужен. Сообщение об успехе будет выводить вызывающий код
	//fmt.Printf("Книга '%s' была выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
	return nil //Книга успешно выдана
}

// ReturnBook возвращает книгу в библиотеку
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

// DisplayReader выводит полную информацию о пользователе
//Этот метод больше не нужен, потому что мы реализовали String() для Reader
/*func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}*/

func (r Reader) String() string {
	status := ""
	if r.IsActive {
		status = "активен"
	} else {
		status = "не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, № %d, пользователь %s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

// Deactivate делает пользователя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}
func (r *Reader) Activate(){
	r.IsActive = true
}

// Library - наша центральная структура-агрегатор
