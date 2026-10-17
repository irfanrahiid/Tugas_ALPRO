//program ini saya buat untuk mengkonversi suhu dari celcius ke kelvin

package main

import "fmt"

func main() {
    var celcius,kelvin float64
    fmt.Println("input celcius =")
    fmt.Scan(&celcius)
    kelvin = celcius + 273
    fmt.Println("hasil =", kelvin)
}