package Goserver

import (
	//"net/http"
	"github.com/joho/godotenv"
	"fmt"

)
func Goserver(){
	 env := godotenv.Load()
	fmt.Println(env)
	fmt.Println("Server")
	//custom http server
	// server_config := &http.Server{
	// Addr: "",	
	// }
}