package main

import (
    "context"
    "fmt"
    "log"
    "time"
	"go-learning/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

func savePerson(ctx context.Context, pool *pgxpool.Pool, person model.Person) (int64, error) {

    var id int64

    err := pool.QueryRow(
        ctx,
        `INSERT INTO users (name, age, city)
         VALUES ($1, $2, $3)
         RETURNING id`,
        person.Name,
        person.Age,
        person.City,
    ).Scan(&id)

    if err != nil {
        return 0, err
    }

    return id, nil
}


func main() {

    connString := "host=/tmp port=5432 dbname=go_learning user=ekaterinapodgornova"

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, connString)

    if err != nil {
        log.Fatal("Не удалось создать подключение:", err)
    }

    defer pool.Close()

    err = pool.Ping(ctx)

    if err != nil {
        log.Fatal("Не удалось подключиться к PostgreSQL:", err)
    }

    fmt.Println("Успешное подключение к PostgreSQL!")

	person := model.Person{
    Name: "Ольга",
    Age:  28,
    City: "Алматы",
	}

	id, err := savePerson(ctx, pool, person)

	if err != nil {
    fmt.Println("Ошибка сохранения:", err)
    return
	}

	fmt.Println("Пользователь сохранён! ID:", id)

}