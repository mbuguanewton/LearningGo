package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// assign key value
	//emails["Bob"] = "bob@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	// declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])


	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
