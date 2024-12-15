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

// то есть handlerfunc преобразует нашу функцию в обработчик 
// чтобы не писать handlerfunc("наша функция")
// мы пишем handlefunc("/", "наша функция")

// фреймворк gin
// для создания роутера
// router := gin.New()
// можно сделать вот так 
// router.GET()
// но чтобы собрать их по группам и построить систему используется gin.Engine

// gin.Engine — это основная структура в фреймворке Gin, которая отвечает за маршрутизацию запросов, обработку middleware, и управление жизненным циклом HTTP-сервера. Внутри gin.Engine содержится множество полей и методов, которые позволяют настраивать и управлять поведением сервера.

// Основные поля и методы gin.Engine:
// Поля:

// RouterGroup: Группа маршрутов, которая содержит все зарегистрированные маршруты.

// trees: Карта деревьев маршрутов, где ключом является HTTP-метод (GET, POST, PUT, DELETE и т.д.).

// pool: Пул контекстов gin.Context, который используется для повторного использования контекстов и экономии памяти.

// UseRawPath: Флаг, указывающий, следует ли использовать необработанный путь запроса.

// UnescapePathValues: Флаг, указывающий, следует ли экранировать значения пути.

// HandleMethodNotAllowed: Флаг, указывающий, следует ли обрабатывать запросы с неправильным методом.

// ForwardedByClientIP: Флаг, указывающий, следует ли использовать заголовок X-Forwarded-For для определения IP-адреса клиента.

// AppEngine: Флаг, указывающий, работает ли сервер на Google App Engine.

// ContextWithFallback: Флаг, указывающий, следует ли использовать резервный контекст.

// Методы:

// Маршрутизация:

// GET(relativePath string, handlers ...HandlerFunc) IRoutes: Регистрирует обработчик для GET-запросов.

// POST(relativePath string, handlers ...HandlerFunc) IRoutes: Регистрирует обработчик для POST-запросов.

// PUT(relativePath string, handlers ...HandlerFunc) IRoutes: Регистрирует обработчик для PUT-запросов.

// DELETE(relativePath string, handlers ...HandlerFunc) IRoutes: Регистрирует обработчик для DELETE-запросов.

// Any(relativePath string, handlers ...HandlerFunc) IRoutes: Регистрирует обработчик для всех HTTP-методов.

// Group(relativePath string, handlers ...HandlerFunc) *RouterGroup: Создает группу маршрутов.

// Middleware:

// Use(middleware ...HandlerFunc) IRoutes: Добавляет middleware в цепочку обработки запросов.

// Запуск сервера: