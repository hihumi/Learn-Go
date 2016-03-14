package main

import "fmt"

func main() {
    var a = 1234567890
    fmt.Println(a)

    var b = 10 * 2
    fmt.Println(b)

    var c, d = 3.1415, "abc"
    fmt.Println(c, d)

    var x int
    x = 1
    fmt.Println(x)

    var y, z int
    y = 2
    z = 3
    fmt.Println(y, z)

    var (
        i int // or // i int = 123
        s1 string // or // s1 string = "aaa"
    )
    i = 123
    fmt.Println(i)
    s1 = "aaa"
    fmt.Println(s1)

    var iScond int
    fmt.Println(iScond) // initial value is 0.

    var e, f, g int = 100, 200, 300
    fmt.Println(e, f, g)

    var (
        h int = 1000
        s2, s3 = "s2s2s2", "s3s3s3"
    )
    fmt.Println(h, s2, s3)

    var iFirst int = 1234567890
    fmt.Println(iFirst)
    var iSecond = iFirst
    fmt.Println(iSecond)

    var bFirst = true
    fmt.Println(bFirst)
    var sFirst = "Golang"
    fmt.Println(sFirst)
    var iThisIsAnInt = 123
    fmt.Println(iThisIsAnInt)
    var fThisIsAFloat = 1.23
    fmt.Println(fThisIsAFloat)
    var im = 1.23i
    fmt.Println(im)
    
    var initialValueFirst = iThisIsAnInt
    fmt.Println(initialValueFirst)
    var initialValueSecond = fThisIsAFloat
    fmt.Println(initialValueSecond)
}

