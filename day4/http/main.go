package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "go-learning/model"
    "log"
)


//функция, которая будет обрабатывать запрос браузера
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//w → ответ клиенту -- куда отправить отчет
	//r → информация о запросе от клиента -- всё, что клиент прислал серверу
    fmt.Fprintln(w, "Привет! Это мой Go backend")
	// w - это инструмент, через который backend отправляет ответ клиенту
	fmt.Fprintln(w, "Метод:", r.Method)
    fmt.Fprintln(w, "Путь:", r.URL.Path)
}

//ctrl+c завершает сервер в терминале

func personHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
    http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
    return
}
//Если запрос пришёл не методом GET — отправить ошибку 405 и выйти из обработчика.

    id := r.URL.Query().Get("id")

    people := map[string]model.Person{
        "1": {
            Name: "Екатерина",
            Age:  29,
            City: "Москва",
        },
        "2": {
            Name: "Анна",
            Age:  26,
            City: "Казань",
        },
    }

//Мы пока не подключали базу данных. Поэтому создаём map прямо в коде, чтобы было на ком тренироваться.
//это временная картотека с двумя вымышленными пользователями

    person, ok := people[id]

    if !ok {
        http.Error(w, "Пользователь не найден", http.StatusNotFound)
        return
    }
//получает пользователя по ID. Переменная ok сообщает, нашёлся ли такой ключ в map.
//Если ok == true, пользователь найден.
//Если ok == false, отправляем ответ 404 Not Found и прекращаем работу обработчика через return.
    
    data, err := json.Marshal(person)

    if err != nil {
        log.Println("Ошибка Marshal:", err)

        http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
        return
    }
    //ошибка подготовки json запишем ее в логи для разработчиков
    //Это отправляет ошибку в HTTP-ответе браузерy:
    //Сообщение отправляется клиенту, а HTTP-ответ получает статус 500 — Internal Server Error.

    w.Header().Set("Content-Type", "application/json; charset=utf-8")// сообщаем, ЧТО отправляем
    w.Write(data)//отправляем сами данные
    //Заголовок нужен, чтобы сообщить браузеру или приложению, в каком формате сервер отправляет ответ,
    //чтобы программы могли правильно взаимодействовать друг с другом. 
}

    //if err != nil {
    //    fmt.Fprintln(w, "Ошибка:", err)
      //  return
    //}
    //Сообщение выводится в Terminal разработчика. Клиенту оно само по себе не отправляется.

    //fmt.Fprintln(w, string(data))


func main() {
    http.HandleFunc("/", helloHandler)
	//«Когда кто-нибудь обратится к /, запусти helloHandler

    http.HandleFunc("/person", personHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")

    http.ListenAndServe(":8080", nil)
	//запускает сервер на порту 8080.

//Запомни: main() запускает сервер и 
//регистрирует маршруты, а personHandler определяет, что именно вернуть при запросе /person.
}
