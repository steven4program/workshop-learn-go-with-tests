package tinymurky

import "fmt"

const (
 englishHello = "Hello"
 spanishHello = "Hola"
 frencHello = "Bonjour"

 spanish = "Spanish"
 french = "French"
)
func Hello(name string, lang string) string {
	

	if name == "" {
		name = "World"
	}

	helloPrefix := generatePrefix(lang)

	return fmt.Sprintf("%s, %s", helloPrefix, name)
}

func generatePrefix(lang string) (prefix string) {

	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frencHello
	default:
		prefix = englishHello
	}
	return
}