package tinhello

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return prefix(language) + name
}

func prefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = "hola, "
	case "french":
		prefix = "bonjour, "
	default:
		prefix = "hello, "
	}
	return
}
