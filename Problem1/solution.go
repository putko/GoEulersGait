package main
import "fmt"
func sumOfAllNumbersTillN(n int) int {
    return n * (n+1) / 2
}

func main() {
    var caseCount = 0
    fmt.Scanln(&caseCount) 
    for i := 0; i < caseCount; i++ {
        var limit int
        fmt.Scanln(&limit)
        limit -= 1
        // Sum all numbers till limit/3 and then multiple it with 3 to get sum of numbers divided by 3. Add the amount you found for 5 then substract result of 15 because it would be dublicate.
        var result = sumOfAllNumbersTillN(limit/3)*3 + sumOfAllNumbersTillN(limit/5)*5 - sumOfAllNumbersTillN(limit/15)*15 
        fmt.Println(result)
    }
}
