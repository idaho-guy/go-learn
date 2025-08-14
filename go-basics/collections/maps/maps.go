package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["IBM"] = "ibm.com"
	fmt.Println(websites)

	fmt.Println(websites["linkedin"])

	delete(websites, "Google")
	fmt.Println(websites)
}
