# interface
an easy way to reuse same code for different type of data
refactor the existing function shuffle to interfaces


type englishBot struct

func (englishBot) getGreeting() string
    return "hello there!"

func printGreeting(eb englishBot)
    fmt.Println(eb.getGreeting())
----

type spanishBot struct

func (spanishBot) getGreeting() string
    return "Hola!"

func printGreeting(sb  spanishBot)
    fmt.Println(sb.getGreeting())



Interfaces are not generic types 
Interfaces are 'implicit'
Interfaces are a contract to help us manage types
Interfaces are tough, Step #1 is understanding how to read them

https://pkg.go.dev/net/http@go1.20.5#Get
https://pkg.go.dev/net/http@go1.20.5#Response

```
func Get(url string) (resp *Response, err error)
```

```
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
...
	Body io.ReadCloser                          ==========>   type ReadCloser interface {
    ...                                                           Reader                   ===============>  type Reader interface {
                                                                  Closer                              	         Read(p []byte) (n int, err error)
                                                                  }                                              }
```

https://pkg.go.dev/io#ReadCloser


```
type Reader interface {
	Read(p []byte) (n int, err error)
}

the thing who want to read the response body, firstly create a byte slice and parse it to reader interface
then let the interface implementor fulfill or fill the data into the byte slice, it doesn't return the byte slice but it let the thing who want to read it can read it.
```


```
type Closer interface {
	Close() error
}
```


```
type error interface {
	Error() string
}
```
