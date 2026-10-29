package main
import "fmt"

//slice (срез) позволяет хранить много значений в одной переменной
//представляют последовательность элементов одного типа переменной длины. 
//В отличие от массивов длина в срезах не фиксирована и динамически может меняться, 
// то есть можно добавлять новые элементы или удалять уже существующие.

func main() {
cities := []string{"Москва", "Казань", "Сочи", "Пермь", "Краснодар"}
fmt.Println(cities[0])
fmt.Println(cities[2])
fmt.Println(cities)
cities = append(cities, "Тюмень")
fmt.Println(cities)
cities = append(cities, "Санкт-Петербург", "Краснодар")
fmt.Println(cities)
fmt.Println("Количество городов:", len(cities))
for i := 0; i < len(cities); i++ {
    fmt.Println(cities[i])
}
for i := 0; i < len(cities); i++ {
    fmt.Println( i+1, cities[i])
}
//другой способ
for i, city := range cities {
    fmt.Println(i+1, city)
}
ages := []int{18, 25, 31, 22, 29}
max := ages[0]
min:= ages[0]
for _, el := range ages {
    //for i, el := range ages; _ , говорит что этот индекс не нужен, а элемент нужен
    if el > max {
        max = el
    }
    if el < min {
        min = el
    }
}
 fmt.Println(max)
  fmt.Println(min)

  for i, city := range cities {
    if city == "Казань" {
        cities[i] = "Тула"
    }
    if city == "Пермь" {
        cities[i] = "Омск"
    }
    if city == "Сочи" {
        cities = append(cities[:i], cities[i+1:]...)
        //cities[:i]       → всё ДО Казани
        // //cities[i+1:]     → всё ПОСЛЕ Казани
        break
    }
  }
  fmt.Println(cities)

  //cities[1]="Тула"
  //cities[3]="Омск"
  //fmt.Println(cities)
//cities[:i] / cities[i+1:]
numbers := []int{4, 15, 7, 20, 3, 12}
for i, number := range numbers {
    if number < 10 {
        numbers = append(numbers[:i], numbers[i+1:]...)
    }
}
 fmt.Println(numbers)
 numbers = []int{4, 15, 7, 20, 3, 12}

var result []int

for _, numberbb := range numbers {
    if numberbb >= 10 {
        result = append(result, numberbb)
    }
}

fmt.Println(result)
}