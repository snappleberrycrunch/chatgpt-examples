package main

import "fmt"

func calculateSER(strikeouts int, inningsPitched float64) float64 {
    // SER = (strikeouts / inningsPitched) * 9
    return (float64(strikeouts) / inningsPitched) * 9
}

func main() {
    strikeouts := 150
    inningsPitched := 180.0

    SER := calculateSER(strikeouts, inningsPitched)
    fmt.Printf("Strikeout Efficiency Rating (SER): %.2f\n", SER)
}

