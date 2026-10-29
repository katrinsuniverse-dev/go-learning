package model

import "fmt"

//`теги` дя json 
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func (p Person) SayHello() {
    fmt.Println("Привет, меня зовут", p.Name)
}

func (p Person) Introduce() {
    fmt.Println("Привет, меня зовут", p.Name, ", мне", p.Age, "лет, я живу в городе", p.City)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

func (p *Person) HaveBirthday() {
    p.Age++
}