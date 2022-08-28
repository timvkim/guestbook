package main

import (
	"fmt"
	"guesbookApp/pkg"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Функция `check` для обработки ошибок
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Объявляем тип структуры для хранения и количества записей
// в качестве параметра `data` в методе `Execute`
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

// Функция-обработчик `viewHandler`
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// Читаем файл `signatures.txt` via `GetStrings` func
	signatures := pkg.GetStrings("signatures.txt")

	// Содержимое view.html используется для создания
	// нового значения Template
	html, err := template.ParseFiles("./view.html")
	check(err)

	// Создаем новую структуру Guestbook
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	// Содержимое шаблона записывается в ResponseWriter
	// с данными из структуры `guestbook`
	err = html.Execute(writer, guestbook)
	check(err)
	// Преобразуем строку в сегмент байтов и добавим в ответ
	//placeholder := []byte("Signature list goes here")
	//_, err := writer.Write(placeholder)
	//check(err)
}

// функция-обработчик для добавления новых записей
func formHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("./form.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// функция-обработчик для отправки форм HTML
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	//_, err := writer.Write([]byte(signature)) - для проверки доступности поля формы
	check(err)
	// Записывает текст в новой строке файла
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	// Перенаправление HTTP
	http.Redirect(writer, request, "/guestbook/", http.StatusFound)
}

func main() {
	log.Printf("Server starting on port %v\n", 8080)
	http.HandleFunc("/guestbook/", viewHandler)
	http.HandleFunc("/guestbook/form", formHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
