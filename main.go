package main

import (
	"fmt"
	"project2/cmd/cli"
	"project2/library"
)

func main() {
    myLibrary := library.New()

    cli.Run(myLibrary)
    
    // err := myLibrary.IssueBookToReader(1, 1)
    // if err != nil {
    //     fmt.Println("Ошибка выдачи:", err)
    // }

    // book, _ := myLibrary.FindBookByID(1)
    // if book != nil {
    //     fmt.Println("Статус книги после выдачи:", book)
    // }

    // err = myLibrary.IssueBookToReader(99, 1)
    // if err != nil {
    //     fmt.Println("Ожидаемая ошибка:", err)
    // }


	// fmt.Println()

	// books := myLibrary.GetAllBooks()
	// for _, book := range books{
	// 	fmt.Println(book)
	// }



    // fmt.Println("Запуск системы управления библиотекой...")

    // myLibrary := &Library{}

    // fmt.Println("\n--- Наполняем библиотеку ---")

    // myLibrary.AddReader("Агунда", "Кокойты")
    // myLibrary.AddReader("Сергей", "Меняйло")

    // myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
    // myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

}
