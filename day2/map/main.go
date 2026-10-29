package main
import "fmt"

//map - представляет ссылку на хеш-таблицу - структуру данных, где каждый элемент представляет пару "ключ-значение". 
// При этом каждый элемент имеет уникальный ключ, по которому можно получить значение элемента.

func main() {    
people := map[string]int{
    "Екатерина": 29,
    "Анна": 25,
    "Мария": 31,
}
fmt.Println(people["Екатерина"])
people["Анна"]=26
people["Ольга"]=28
delete(people,"Мария")
fmt.Println(people)

for key, value := range people {
    if value >= 28 {
        fmt.Println(key,value)
    }
}
//порядок в range для map может рандомным быть
}
