pkg main
import (
    "fmt"
    "time"
)
 
func say() {
    fmt.Println("testing goroutine baru")
}
func main() {
    go say()
    time.Sleep(1 * time.Second)//jika ingin menggunakan time.Sleep
    fmt.Println("Fungsi Utama")
}
//gye wong
