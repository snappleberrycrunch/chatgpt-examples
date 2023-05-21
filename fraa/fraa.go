package main

import "fmt"

func calculateFRAA(putouts, assists, errors, leaguePutouts, leagueAssists, leagueErrors float64) float64 {
    // FRAA = (putouts - leaguePutouts) + (assists - leagueAssists) - (errors - leagueErrors)
    return (putouts - leaguePutouts) + (assists - leagueAssists) - (errors - leagueErrors)
}

func main() {
    putouts := 120.0
    assists := 50.0
    errors := 5.0
    leaguePutouts := 100.0
    leagueAssists := 40.0
    leagueErrors := 10.0

    FRAA := calculateFRAA(putouts, assists, errors, leaguePutouts, leagueAssists, leagueErrors)
    fmt.Printf("Fielding Runs Above Average (FRAA): %.2f\n", FRAA)
}
