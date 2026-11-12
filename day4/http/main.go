package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "go-learning/model"
    "log"
    "strings"
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

//Первая функция — получение пользователя
func personHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
    http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
    return
}
//Если запрос пришёл не методом GET — отправить ошибку 405 и выйти из обработчика.
//r.Method — метод запроса, который пришёл от клиента
//http.MethodGet — константа, обозначающая метод GET.

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


    //POST функция создания пользователя

    func createPersonHandler( w http.ResponseWriter, r *http.Request){
        if r.Method != http.MethodPost {
            http.Error (w, "Метод не поддерживается", http.StatusMethodNotAllowed)
            return
        }

        var person model.Person

        err := json.NewDecoder(r.Body).Decode(&person)
        //r.body Получает тело HTTP-запроса
        //json.NewDecoder Создаёт инструмент для чтения JSON
        //.decode преобразует json в person 
        //JSON читается непосредственно из HTTP-запроса И снова используется &person, 
        // потому что мы хотим заполнить нашу переменную person.

        if err != nil {
            http.Error (w, "Некорректный Json", http.StatusBadRequest)
            return
        }

        if strings.TrimSpace(person.Name) == "" {
            http.Error(w, "Ваше Имя обязательно", http.StatusBadRequest)
            return
        }
        //strings.TrimSpace() убирает пробелы в начале и конце строки.
        //== "" проверяет, осталась ли строка пустой.
        // http.StatusBadRequest — статус 400: клиент отправил некорректные данные.

        if person.Age < 0 {
            http.Error (w, "Возраст не может быть отрицательным", http.StatusBadRequest)
            return
        }
    
        data,err := json.Marshal(person)
        
        if err != nil {
            http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
        return 
        }
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)
        //json.NewEncoder(w).Encode(person)
        //Он сразу преобразует person в JSON и записывает его в HTTP-ответ
    }

func main() {
    http.HandleFunc("/", helloHandler)
	//«Когда кто-нибудь обратится к /, запусти helloHandler

    http.HandleFunc("/person", personHandler)
    
    http.HandleFunc("/people", createPersonHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")

    http.ListenAndServe(":8080", nil)
	//запускает сервер на порту 8080.

//Запомни: main() запускает сервер и 
//регистрирует маршруты, а personHandler определяет, что именно вернуть при запросе /person.
}
