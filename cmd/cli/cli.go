package cli

import (
	"bufio"
	"fmt"
	"project2/library"
	"project2/storage"
	"os"
	"strconv"
)

func printMenu() {

	fmt.Println("--SIMPLE-LIBRARY--")
	fmt.Println("----------------------------------------")
	fmt.Println("-MENU-")
	fmt.Println("[1] Добавление книги")
	fmt.Println("[2] Выдача книги")
	fmt.Println("[3] Возврат книги")
	fmt.Println("[4] Поиск книги")
	fmt.Println("[5] Импорт книги")
	fmt.Println("[6] Экспорт книг")
	fmt.Println("[7] Вывод списка книг")
	fmt.Println("[8] Добавление читателя")
	fmt.Println("[9] Поиск читателя")
	fmt.Println("[10] Импорт читателя")
	fmt.Println("[11] Экспорт читателя")
	fmt.Println("[0] Выход")

}

func handlerChoice(choice int, scanner *bufio.Scanner, library *library.Library) {
	switch choice {
	case 1:
		fmt.Println("Введите название книги: ")
		scanner.Scan()
		title := scanner.Text()

		fmt.Println("Введите автора книги: ")
		scanner.Scan()
		author := scanner.Text()

		fmt.Println("Введите название книги: ")
		scanner.Scan()
		year, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Год должен состоять из цифр")
			return
		}

		if _, err := library.AddBook(title, author, year); err != nil {
			fmt.Printf("Произошла ошибка при добавление книги:%v", err)
		}

		// выдача книг читателю
	case 2:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}

		fmt.Println("Введите номер Читателя: ")
		scanner.Scan()
		idUser, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		if err := library.IssueBookToReader(idBook, idUser); err != nil {
			fmt.Printf("Произошла ошибка при выдаче книги:%v", err)
		}

	case 3:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		err = library.ReturnBook(idBook)
		if err != nil{
			fmt.Println("Ошибка возврата книги ", err)
			return
		}
		fmt.Println("Книга возвращена")

	case 4:
		fmt.Println("Введите номер книги: ")
		scanner.Scan()
		idBook, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Номер должен состоять из цифр!")
			return
		}
		book, err := library.FindBookByID(idBook)
		if err != nil{
			fmt.Println("Ошибка при поиске книги", err)
			return
		}
		fmt.Println(book)

	case 5:
		fmt.Println("Введите путь к файлу: ")
		scanner.Scan()
		filename := scanner.Text()
		books, err := storage.LoadBooksFromCSV(filename)
		if err != nil{
			fmt.Println("Ошибка экспорта: ", err)
			return
		}
		library.Books = books
		fmt.Println("Успешно импортировано")
	case 6:
		fmt.Println("Введите название файла:")
		scanner.Scan()
		filename := scanner.Text()
		err := storage.SaveBooksToCSV(filename,library.Books)
		if err != nil{
			fmt.Println("Ошибка экспорта: ")
			return
		}
		fmt.Println("Успешный экспорт")
	case 7:
		books := library.GetAllBooks()
		fmt.Println("Список книг")
		for _, book := range books{
			fmt.Println(book)
		}
	case 8:
		fmt.Println("Введите имя: ")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Println("Введите фамилию: ")
		scanner.Scan()
		lastName := scanner.Text()
		_, err := library.AddReader(firstName,lastName) 
		if err != nil{
			fmt.Println("Ошибка добавления: ", err)
			return
		}
		fmt.Println("Пользователь успешно добавлен")
	// case 9:
	// 	fmt.Println("Введите ID читателя: ")
	// 	scanner.Scan()
	// 	userID := scanner.Text(Atoi())
	case 10:
		fmt.Println("Введите название файла:")
		scanner.Scan()
		filename := scanner.Text()
		readers, err := storage.LoadReaderFromCSV(filename)
		if err != nil{
			fmt.Printf("Не удалось получить данные: ", err)
			return
		}
		fmt.Println(readers)
	case 11:
		fmt.Println("Введите название файла:")
		scanner.Scan()
		filename := scanner.Text()
		err := storage.SaveReaderToCSV(filename, library.Readers)
		if err != nil{
			fmt.Println("Ошибка ", err)
			return
		}
		fmt.Println("Succes")
	case 12:
		readers := library.GetAllReaders()
		fmt.Println("Reader list: ")
		for _, reader := range readers{
			fmt.Println(reader)
		}
	case 0:
		fmt.Println("Bye")
	}

}

func Run(lib *library.Library){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil{
			fmt.Println(err)
			continue
		}
		handlerChoice(choice,scanner,lib)
		if choice == 0{
			storage.SaveBooksToCSV("books.csv", lib.Books)
			if err != nil{
				fmt.Println("Произошла ошибка экспорта: ", err)
			}
			break
		}
	}
}