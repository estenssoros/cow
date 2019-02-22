package main

func main() {
	RunServer()
}

func RunServer() {
	app := NewApp()
	app.Start(":3001")
}
