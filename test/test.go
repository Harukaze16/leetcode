package main
import "fmt"

type AS struct{
    Var int
    XX int
}

func main() {
    var p *AS= nil
    if p {
        fmt.Println("no")
    } else {
        fmt.Println("yes")
    }
}
