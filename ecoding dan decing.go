package main
 
import (
    "encoding/json"
    "fmt"
    "log"
)
 
type Mahasiswa struct {
    NamaDepan    string `json:"nama_depan"`
    NamaBelakang string `json:"nama_belakang"`
}
 
func encodeJSON(data []Mahasiswa) []byte {
    dataJSON, err := json.Marshal(data)
 
    if err != nil {
        log.Fatal(err)
    }
 
    return dataJSON
}
 
func main() {
 
    mhs := []Mahasiswa{{
        NamaDepan:    "gye",
        NamaBelakang: "wong",
    },
        {
            NamaDepan:    "vha",
            NamaBelakang: "wang",
        },
    }
 
    dataJSON := encodeJSON(mhs)
 
    fmt.Println(string(dataJSON))
}
