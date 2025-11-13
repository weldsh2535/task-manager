package main

func main() {
	InitDB()
	r := SetupRouter()
	r.Run(":8080")
}
