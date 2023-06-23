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