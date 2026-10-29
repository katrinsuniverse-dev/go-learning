//циклы
package main
import "fmt"
func main() {
var sum = 0
for i := 1; i <= 20; i++ {
    if i % 2 == 0 {
		fmt.Println(i)
		sum = sum + i
	}
}
fmt.Println(sum)
}