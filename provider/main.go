package main

import "log"

func main() {
	user := User{}
	info, err := user.ProcessInfo()
	log.Println(info, err)
}
