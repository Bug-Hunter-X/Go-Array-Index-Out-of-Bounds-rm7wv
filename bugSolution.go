func main() {
    var a [10]int
    if len(a) > 5 {
        fmt.Println(a[5])
    } else {
        fmt.Println("Index out of bounds")
    }
}