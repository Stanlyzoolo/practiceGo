package databases

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Book struct {
	title    string	
	price    string	
	quantity string	
}

func ReadingCSV(csvName string) {
	csvFile, error := os.Open(csvName)
	if error != nil {
		panic(error)
	}

	reader := csv.NewReader(csvFile)
	fieldsStruct, _ := reader.ReadAll()
	
	for _, field := range fieldsStruct {
		book := Book {
			title: field[0],
			price: field[1],
			quantity: field[2],
		}
		fmt.Printf("The book named `%s` costs %s рублей Р.Б. The quantity is %s\n", book.title, book.price, book.quantity)
	}

}
