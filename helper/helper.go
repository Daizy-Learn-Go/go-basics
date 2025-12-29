package helper

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Goodbye " + name
}

func Contoh(name string) string {
	return sayGoodBye(name)
}

func SayHello(name string) string {
	return "Hello " + name
}
