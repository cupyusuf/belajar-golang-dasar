package helper

var version = "1.0.0" // private (tidak bisa diakses diluar package)
var Application = "Golang"

// Tidak bisa diakses diluar package
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
