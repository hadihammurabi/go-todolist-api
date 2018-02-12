package main

import "TodoAPI/server"

func main() {
	server.Listen(Cfg("host"))
}
