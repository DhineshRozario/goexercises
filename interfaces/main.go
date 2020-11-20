package main

import "fmt"


type englishBot struct {}

type spanishBot struct {}



func main() {

}

func (eb englishBot) getGreeting() string {
    return "Hello there!"
}

func (sp spanishBot) getGreeting() string {
    return "Hola!"
}

func printGreeting(eb englishBot) {
    fmt.Println(eb.getGreeting())
}

func printGreeting(sp spanishBot) {
    fmt.Println(sp.getGreeting())
}