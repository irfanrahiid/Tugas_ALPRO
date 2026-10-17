//program ini saya buat untuk mengconversi suhu dari celcius ke fahrenheit

package main

import "fmt"

func main() {
    var celcius,fahrenheit float64
    fmt.Println("input celcius =")
    fmt.Scan(&celcius)
    fahrenheit = 9.0 / 5.0 * celcius + 32.0
    fmt.Println("hasil =")
    fmt.Println(fahrenheit)
}