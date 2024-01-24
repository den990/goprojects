package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type TODO struct {
	ID    int
	Title string
}

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`)

	if err != nil {
		log.Fatal(err)
	}

	var choice int

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1.Добавить запись")
		fmt.Println("2.Удалить задачу")
		fmt.Println("3.Просмотреть задачи")
		fmt.Println("4.Выход")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addTodo(db)
		case 2:
			removeTodo(db)
		case 3:
			showTodo(db)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Ошибка! Выберите правильно действие")

		}
	}
}

func addTodo(db *sql.DB) {
	var title string
	fmt.Println("Введите название задачи:")
	title, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	_, err := db.Exec("INSERT INTO todos (title) VALUES (?)", title)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Задача добавлена успешно.")
}

func showTodo(db *sql.DB) {

	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Список задач:")
	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d. %s\n", todo.ID, todo.Title)
	}
}

func removeTodo(db *sql.DB) {
	var id int
	fmt.Print("Введите ID задачи для удаления: ")
	fmt.Scanln(&id)
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Задача удалена успешно.")
}
