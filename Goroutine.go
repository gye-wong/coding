//goroutine type 1
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

//goroutine type2
pkg main
import (
    "fmt"
    "time"
)
 
func cetakAngka() {
    angka := []int{1, 2, 3, 4, 5}
    for x := range angka {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(angka[x])
    }
}
 
func cetakHuruf() {
    huruf := []string{"a", "b", "c", "d", "e"}
    for x := range huruf {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(huruf[x])
    }
}
 
func main() {
    go cetakAngka()
    go cetakHuruf()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("Fungsi Utama")
}
