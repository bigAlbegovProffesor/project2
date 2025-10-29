package storage

import (
	"fmt"
	"log"
	"os"
	"project2/domain"
)

type Storable interface {
	Save() error
	Load() error
}

func Save(filename string, book []*domain.Book) {
	fmt.Println("This goes into the text file")
	data := []byte("first line\n second line")
	err := os.WriteFile("log.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File log.txt created")

}
