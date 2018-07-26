package greeter

import "fmt"

// EnglishGreeter is a greeter that will return greetings in English.
type EnglishGreeter struct {
	// You could provide some more configuration to something like this too, maybe to replace the
	// word "Hello"?
}

// NewEnglishGreeter returns a new EnglishGreeter instance.
func NewEnglishGreeter() *EnglishGreeter {
	return &EnglishGreeter{}
}

// Greet provides an English greeting.
func (g *EnglishGreeter) Greet(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
