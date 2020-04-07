package database

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestDatabaseConnection(t *testing.T) {
	//创建对象，不会创建实际的连接
	pool, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(10)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	//接受系统层面的信号量
	appSignal := make(chan os.Signal, 3)

	//监听停止信号
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		//直接阻塞在这个信号即可，不需要使用for循环
		select {
		//如果发过来了停止信号，那么停止
		case <-appSignal:
			stop()
		}
	}()

	Ping(ctx, pool)

	Insert(ctx, pool)

	Query(ctx, pool)
}

//
func Insert(ctx context.Context, pool *sql.DB) {

	result, err := pool.ExecContext(ctx,
		"INSERT INTO people (name) values (?)",
		"suitingwei",
	)

	if err != nil {
		log.Fatalln("Failed to insert data into people table:", err.Error())
		return

	}
	id, err := result.LastInsertId()

	log.Printf("Insert successfully, result is %d\n", id)
}

func Query(ctx context.Context, pool *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := pool.QueryContext(ctx, "select id,name from people ")

	if err != nil {
		log.Fatal("unable to execute search query ", err)
	}

	var users []User

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	jsonUsers, _ := json.Marshal(users)

	log.Printf("Current users:%s\n", jsonUsers)
}

func Ping(ctx context.Context, pool *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalln("Failed to connect to the database", err.Error())
	}
}
