package main

func main() {
	gin_router := setup_router()
	listenPort := "8086"
	// Listen and Server on the LocalHost:Port
	gin_router.Run(":" + listenPort)
}
