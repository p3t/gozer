package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the goZer!\n")
	cmd := exec.Command("ssh", "localhost:2222", "ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, string(out[:]))
	}
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
