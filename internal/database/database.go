package database

import (
	"context"
	"log"
	"todo/internal/models"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Conn *pgx.Conn
}

func New(dbConf models.DBConfig) *Connection {
	connectionString := "postgres://" + dbConf.Username + ":" + dbConf.Password + "@" + dbConf.HostName + ":" + dbConf.DBPort + "/" + dbConf.DBName
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Error while connecting database: %s", err)
	}

	return &Connection{
		Conn: conn,
	}
}

func (conn *Connection) ReadToDo(idUser string) []string {

	var resultQuery []string

	result, err := conn.Conn.Query(context.Background(),
		"SELECT * FROM tasks where iduser=$1", idUser)
	if err != nil {
		log.Fatalf("Error while exec sql-query: %s, %s", err, "SELECT * FROM tasks WHERE iduser=?")
	}

	for result.Next() {
		var ToDoObject models.ToDoObject
		err := result.Scan(&ToDoObject.Id, &ToDoObject.TodoText, &ToDoObject.IdUser)
		log.Println()
		if err != nil {
			log.Fatalf("Error while scanning querry result: %s", err)
		}

		resultQuery = append(resultQuery, ToDoObject.TodoText)
	}

	log.Println(resultQuery)
	return resultQuery
}

func (conn *Connection) DeleteTask(taskId string) {
	_, err := conn.Conn.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", taskId)
	if err != nil {
		log.Fatalf("Error while deleting task: %s", err)
	}
}

func (conn *Connection) AddTask(task, idUser string) {
	_, err := conn.Conn.Exec(context.Background(), `INSERT INTO tasks (task, iduser) VALUES ($1, $2)`, task, idUser)
	if err != nil {
		log.Fatalf("Error while adding new task: %s", err)
	}
}

func (conn *Connection) UpdateTask(task, taskId string) {
	_, err := conn.Conn.Exec(context.Background(), `
		UPDATE tasks
		SET task=$1
		WHERE id=$2
	`, task, taskId)
	if err != nil {
		log.Fatalf("Error while updating task: %s", err)
	}
}
