package main

func main() {
	server := NewAPIServer(":7000")

	server.Run()
}
