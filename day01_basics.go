package main
import "fmt"
func main() {
	fmt.Println("Привет!")
	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scan(&name)
	//& - Вот адрес переменной name. Запиши введённое значение именно туда
	fmt.Println("Приятно познакомиться,", name , "!")
	var year int
	fmt.Println ("Какого ты года рождения?")
	fmt.Scan(&year)
	age:= 2026 - year
	var city string
	fmt.Println ("Где ты живешь?")
	fmt.Scan(&city)	
	age5 := age + 5
    var height float64
	fmt.Println("Какой у тебя рост в метрах?")
	fmt.Scan(&height)
	var weight int
	fmt.Println("Какой у тебя вес?")
	fmt.Scan(&weight)
    imt:= float64(weight) / (height * height)
	fmt.Println("Меня зовут", name , "мне", age, ", живу в городе", city,".")
	fmt.Println("Сейчас я изучаю Go")
	fmt.Println("Через 5 лет мне будет", age5)
	fmt.Println("Индекс моей массы тела", imt)

	if age < 18 {
		fmt.Println ("Ты несовершеннолетняя")
	} else if age <= 30 { 
		fmt.Println ("Ты совершеннолетняя и до 30 ещё есть время :)")
		} else {
			fmt.Println ("Ты уже перешагнула 30-летний рубеж")
		}	
}