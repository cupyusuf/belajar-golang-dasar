package main

type blackList func(string) bool

func registerUser(name string, blackList blackList) {
	if blackList(name) {
		println("You are blocked", name)
	} else {
		println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Yusuf", blackList)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
