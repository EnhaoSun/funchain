# funchain
print your function-call chain in console

## colorful output
example test code from [7-days-golang/go-web](https://geektutu.com/post/gee-day3.html)

add route
![image](./funchain1.PNG)

-----------
http "GET" request
![image](./funchain2.PNG)

## usage
add require in go.mod
```go
require github.com/EnhaoSun/funchain v1.0.0 // indirect
```
-----------
add defer at the beginning of your function
```go
import "github.com/EnhaoSun/funchain"

func () xxx() {
    defer funchain.Trace()()
    //....
}
```