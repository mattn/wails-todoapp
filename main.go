package main

import (
	"database/sql"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/leaanthony/mewn"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails"
)

type Todo struct {
	Id        int64      `json:"id" db:"id"`
	Body      string     `json:"body" db:"body"`
	Done      bool       `json:"done" db:"done"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type TodoApp struct {
	dbmap *gorp.DbMap
}

func NewTodoApp() *TodoApp {
	return &TodoApp{}
}

func (app *TodoApp) WailsInit(runtime *wails.Runtime) error {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	table := dbmap.AddTableWithName(Todo{}, "items")
	table.SetKeys(true, "id")
	table.ColMap("CreatedAt").DefaultValue = "(datetime('now', 'localtime'))"
	table.ColMap("UpdatedAt").DefaultValue = "(datetime('now', 'localtime'))"
	app.dbmap = dbmap
	return nil
}

func (app *TodoApp) Add(body string) error {
	return app.dbmap.Insert(&Todo{Body: body})
}

func (app *TodoApp) Done(id int64, done bool) error {
	_, err := app.dbmap.Exec("update items set done = ? where id = ?", done, id)
	return err
}

func (app *TodoApp) List() ([]Todo, error) {
	var items []Todo
	_, err := app.dbmap.Select(&items, "select * from items order by created_at desc")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Todo App",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(NewTodoApp())
	app.Run()
}
