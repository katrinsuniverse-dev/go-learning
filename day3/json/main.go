package main
import (
    "encoding/json"
    "fmt"
    "go-learning/model"
    "errors"
)

func findPerson(name string, people []model.Person) (model.Person, error){
    for _, person := range people {
        if person.Name == name {
            return person, nil
        }
    }
    return model.Person{}, errors.New("Такого человека нет")
}

func main() {
	person1 := model.Person{
    Name: "Екатерина",
    Age:  29,
    City: "Москва",
}


//сама функция json.Marshal или json.Unmarshal возвращает nil, когда всё прошло успешно

//MARK: Person -> JSON
//Marshal — это функция, которая преобразует Go-значение в JSON
data, err := json.Marshal(person1) //То есть Marshal создаёт JSON в виде []byte

if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println(string(data))
        //Но есть один важный нюанс: Marshal возвращает не string, а []byte
        //data содержит байты, которые представляют текст
    }

    
//MARK: JSON → Person
 
var personBack model.Person

err = json.Unmarshal(data, &personBack)
//«Возьми эти байты, разбери JSON и запиши полученные данные в personBack»

if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println(personBack)
}

people := []model.Person {
    {Name: "Мария", Age: 30, City: "Питер"},
    {Name: "Анна", Age: 26, City: "Казань"},
    {Name: "Andrew", Age: 20, City: "Kanzas"},
}

person, err := findPerson("Мария", people)

if err != nil {
    fmt.Println("Ошибка:", err)
    } else {
        personJson, err := json.Marshal(person)
        if err != nil {
            fmt.Println("Ошибка:", err)
            } else {
                fmt.Println(string(personJson))
            }
    }
}
