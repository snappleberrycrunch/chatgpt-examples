package main

import "fmt"

func calculateDER(putouts, assists, totalChances float64) float64 {
    // DER = (putouts + assists) / totalChances
    return (putouts + assists) / totalChances
}

func main() {
    putouts := 120.0
    assists := 50.0
    totalChances := 200.0

    DER := calculateDER(putouts, assists, totalChances)
    fmt.Printf("Defensive Efficiency Rating (DER): %.3f\n", DER)
}
