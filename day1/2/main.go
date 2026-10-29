package main
import "fmt"
func sayHello(name string) {
    fmt.Println("Привет,", name) // name — параметр
}
func sayGoodbye() {
	fmt.Println("Доброй ночи")
}
func add(a int, b int) int {
    return a + b
}
func multiply( a int, b int) int {
	return a * b
}

//func isAdult(age int) bool {
//	if age >= 18 {
//		return true
//	} else {
//		return false
//	}
//}

func isAdult(age int) bool {
    return age >= 18
}

func canDrive (age int) bool {
	return age >=18
}

//slice позволяет хранить много значений в одной переменной



func main() {
   sayHello("Екатерина") // "Екатерина" —аргумент
	sayGoodbye()
	result := add(5, 3) //«Создай новую переменную result и положи в неё результат функции»
    fmt.Println(result)
	result = multiply(6,7) //Возьми существующую переменную result и положи в неё результат multiply
	fmt.Println(result)
	adult := isAdult(29)
	fmt.Println(adult)
	adult = isAdult(15)
	fmt.Println(adult)
	welcome := isAdult(29)
	if welcome {
		fmt.Println("Добро пожаловать!")
	} else {
		fmt.Println ("Доступ запрещен")
	}
	drive := canDrive(29)
	if drive {
		fmt.Println("Можно водить")
	} else {
		fmt.Println ("Водить пока нельзя")
	}
	for i := 15; i <= 25; i++ {
		driveI := canDrive(i) //можно не создавать а сразу в иф кинуть canDrive(i)
		if driveI {
			fmt.Println( i, "- Можно водить")
			} else {
		fmt.Println (i, "- Водить пока нельзя")
	}}
}