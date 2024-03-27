package airportrobot

import "fmt"

// Greeter interface
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// German config
type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return "Hallo " + name + "!"
}

// Italian config
type Italian struct{}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

// Portuguese config
type Portuguese struct{}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}

func (g Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

// Hello function
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}
