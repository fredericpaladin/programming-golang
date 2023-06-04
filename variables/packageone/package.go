package packageone

var privateVar = "I'm private"        // this is private because it starts lowercase
var PublicVar = "I'm public/exported" // This is public because it starts with a capital letter!

func privateFunc() {

}

// This is public cause it starts with uppercase letter
func PublicFunc() {

}
