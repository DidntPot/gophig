## Getting Started

Gophig may be imported using `go get`:
```
go get github.com/RestartFU/gophig
```

## Usage

You may create a new `*Gophig`:
```go
type foo struct{
   foo string `toml:"foo"`
   bar string `toml:"bar"`
}

myFoo := &foo{foo: "foo", bar: "bar"}

g := gophig.NewGophig("./config", "toml", 0777)
```
Then you may use the method `SetConf(v interface{})`:
```go
if err := g.SetConf(foo);err != nil{
   log.Fatalln(err)
}

// Output file content:
// ./config.toml
/* 
   foo = "foo"
   bar = "bar"
*/
```
Or the method `GetConf(v interface{})`:
```go
// If we assume that the output file content is the same as the example up there:
var myFooStruct *foo
if err := g.GetConf(foo);err != nil{
   log.Fatalln(err)
}
log.Println(*foo)

// Output:
/* 
   {foo: "foo", bar: "bar"}
*/
```
