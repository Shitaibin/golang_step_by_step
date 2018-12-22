import (
    "fmt"
    "time"
    "math/rand"
)

func eat() chan string {
    out := make(chan string)
    go func (){
        rand.Seed(time.Now().UnixNano())
        time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
        out <- "Mom call you eating"
        close(out)
    }()
    return out
}


func main() {
    eatCh := eat()
    sleep := time.NewTimer(time.Second * 3)
    select {
    case s := <-eatCh:
        fmt.Println(s)
    case <- sleep.C:
        fmt.Println("Time to sleep")
    default:
        fmt.Println("Beat DouDou")
    }
}
