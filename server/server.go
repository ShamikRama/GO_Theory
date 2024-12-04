package main

import (
	"fmt"
	"net/http"
)

// Обработчик для GET-запроса на корневой путь "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Обработчик для GET-запроса на путь "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
}

func main() {
	// Регистрируем обработчик для маршрута "/" с использованием http.HandleFunc
	http.HandleFunc("/", helloHandler)

	// Регистрируем обработчик для маршрута "/about" с использованием http.Handle
	http.Handle("/about", http.HandlerFunc(aboutHandler))

	// Запускаем сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

то есть handlerfunc преобразует нашу функцию в обработчик 
чтобы не писать handlerfunc("наша функция")
мы пишем handlefunc("/", "наша функция")

фреймворк gin
для создания роутера
router := gin.