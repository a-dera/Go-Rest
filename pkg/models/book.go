package models

import (
	"https://github.com/a-dera/Go-Rest/pkg/config"
	"https://github.com/go-gorm/gorm"
)

var database *gorm.DB

type Book struct {
	gorm.Model
	//Id          string `json:"id"`
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	database = config.GetDB()
	database.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	database.NewRecord(b)
	database.Create(&b)
	return b
}

func  GetAllBooks() []Book {
	var Books []Book
	database.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book , *gorm.DB){
	var getBook Book
	db:=database.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	database.Where("ID = ?", ID).Delete(book)
	return book
}