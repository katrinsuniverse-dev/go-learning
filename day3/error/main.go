package main
import (
    "errors"
    "fmt"
    "go-learning/model"
)

//MARK: ошибки error
func divide(a float64, b float64) (float64, error) {
//Она будет возвращать два значения одновременно
//результат + ошибка
	if b == 0 {
        return 0, errors.New("нельзя делить на ноль")
		//error — это тип, а errors.New(...) — функция, которая создаёт конкретное значение этого типа.
    } 
	return a / b, nil
}

//Она должна:
//найти человека по имени; если нашла → вернуть его;если не нашла → вернуть ошибку.
func findPerson(name string, people []model.Person) (model.Person, error) {
	//сейчас people ты создаёшь внутри main, а функция findPerson
	//  находится снаружи main и не имеет доступа к этой локальной переменной.
	for _, person := range people{
		if person.Name == name {
			return person, nil
		}
	}
	return model.Person{}, errors.New("такого человека нет")
	//model.Person{} — это пустой Person
}


func main () {
result, err := divide(10, 5)

if err != nil {
    fmt.Println("Ошибка:", err)
} else {
    fmt.Println("Результат:", result)
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
        fmt.Println(person)
    }
}