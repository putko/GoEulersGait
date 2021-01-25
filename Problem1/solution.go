package main
import "fmt"
func sum(n int) int {
    return n * (n+1) / 2
}

func main() {
	var caseCount = 0
    fmt.Scanln(&caseCount) 
	for i := 0; i < caseCount; i++ {
		var limit int
		fmt.Scanln(&limit)
		limit -= 1
		fmt.Println(sum(limit/3)*3 + sum(limit/5)*5 - sum(limit/15)*15)
	}
}
