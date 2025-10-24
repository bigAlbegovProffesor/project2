package library

import(
	"project2/domain"
	"strings"
	"fmt"
	"errors"
)

type Library struct {
	Books   []*domain.Book
	Readers []*domain.Reader

	//Счетчики для генерации уникальных ID
	lastBookID   int
	lastReaderID int
}

func (lib *Library) AddReader(firstName, lastName string) (*domain.Reader, error){
	cleanedFirstName := strings.ToLower(strings.TrimSpace(firstName))
	cleanedLastName := strings.ToLower(strings.TrimSpace(lastName))
	if (cleanedFirstName == "" || cleanedLastName == ""){
		return nil, errors.New("Фамилия и имя не могут быть пустыми")
	}
	lib.lastReaderID++

	//Создаем нового читателя
	newReader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true, //Новый читатель всегда активный
	}

	//Добавляем читателя в срез
	lib.Readers = append(lib.Readers, newReader)

	return newReader, nil
}

// AddBook добавляет новую книгу в библиотеку
func (lib *Library) AddBook(title, author string, year int) *domain.Book {
	lib.lastBookID++

	//Создаем новую книгу
	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false, //Новая книга всегда в наличии
	}

	//Добавляем новую книгу в библиотеку
	lib.Books = append(lib.Books, newBook)

	fmt.Printf("Добавлена новая книга: %s\n", newBook)
	return newBook
}

// FindBookByID ищет книгу по ее уникальному ID
func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}

	return nil, fmt.Errorf("книга с ID %d не найдена в библиотеке", id)
}

// FindReaderByID ищет читателя по его уникальному ID
func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}

	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// IssueBookToReader - основной публичный метод для выдачи книги
func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	//1. Найти книгу
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	//2. Найти читателя
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	//Вызываем обновленный метод и ПРОВЕРЯЕМ ОШИБКУ
	err = book.IssueBook(reader)
	if err != nil {
		return err
	}     
	return nil //Все 3 шага прошли успешно
}

func (lib *Library) ReturnBook(bookID int) error {

	book, err := lib.FindBookByID(bookID)
	if err != nil{
		return err
	} 
	err = book.ReturnBook()
	if err != nil{
		return err
	}
	book.IsIssued = false
	return nil
}

// ListAllBooksПоказывает все книги в библиотеке
func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}