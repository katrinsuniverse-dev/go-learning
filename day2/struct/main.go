package main
import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

//MARK: Cоздаем методы!

func (p Person) SayHello() {
    fmt.Println("Привет, меня зовут", p.Name)
}
//метод Introduce принадлежит типу Person, а p — это конкретный Person, с которым мы сейчас работаем
func (p Person) Introduce() {
    fmt.Println("Привет, меня зовут", p.Name, ", мне", p.Age, "лет, я живу в городе", p.City)
}

//func Introduce(person Person)
//функция Introduce принимает переменную person, которая имеет тип Person. (обычная функция с параметром)

func (p Person) IsAdult() bool {
   return p.Age >= 18
}
//пишем метод с большой буквы - так к нему можно обратиться из другого package


//MARK: УКАЗАТЕЛИ: 
// 1) * указывает, где находятся данные
func (p *Person) HaveBirthday() {
    p.Age++
}
//*Person означает, что p связан с исходным объектом, поэтому мы можем изменить его Age

// 2) & — взять адрес
//&person - дай мне адрес person
//person := Person{Age: 30}
//pointer := &person

//Здесь Name, Age, City — заранее определённые поля типа Person

//MARK: Вызов func main
func main () {
person := Person {
	Name: "Екатерина",
	Age: 29,
	City: "Москва", 
}
fmt.Println(person.Name)
fmt.Println(person.Age)
fmt.Println(person.City)

//MARK: делаем SLICE из Person
//  из элементов Person, с типом struct
people := []Person{
    {
        Name: "Мария",
        Age:  30,
		City: "Питер",
    },
    {
        Name: "Анна",
        Age:  26,
        City: "Казань",
    },
    {
        Name: "Andrew",
        Age: 20,
        City: "Kanzas",
    },
    {
        Name: "Mark",
        Age: 19,
        City: "NY",
    },
}
fmt.Println(people)

for _, person := range people {
    if person.Age >= 25 {
        fmt.Println(person)
    }
}

oldest := people[0]

for _, person := range people {
    if person.Age > oldest.Age {
        oldest = person
    }
}
//В итоге oldest хранит целый Person, а не только возраст.
fmt.Println(oldest)

person.SayHello()

people[0].SayHello()

person.Introduce()

if people[3].IsAdult() {
    fmt.Println("Совершеннолетний")
    } else {
    fmt.Println("Несовершеннолетний")
}

if person.IsAdult() {
    fmt.Println("Совершеннолетний")
    } else {
    fmt.Println("Несовершеннолетний")
}

people[0].HaveBirthday()
people[1].HaveBirthday()
fmt.Println(people[0])
fmt.Println(people[1])


//& — взять адрес
//&person - дай мне адрес person
//person := Person{Age: 30}
//pointer := &person

}