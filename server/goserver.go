package goserver

import (
	"log"
	"net/http"
	"time"

	"githhub.com/AJ-Brown-InTech/go-server/env"
)
var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)
type dotenv struct{
	port string
	host string
}
type timeHandler struct {
    zone *time.Location
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().In(th.zone).Format(time.RFC1123)
    w.Write([]byte("Current Server Time: " + tm))
}

func newTimeHandler(name string) *timeHandler {
    return &timeHandler{zone: time.FixedZone(name, 0)}
}

func Goserver(){
	//use type to bring in env vars and retrieve for use
	vars := new(dotenv)
	vars.host, vars.port = "HOST", "PORT"
	port := env.GoEnvVariables(vars.port)
	//host := env.GoEnvVariables(vars.host)
	timeout := 10 * time.Second
	//Get local 
	timeZone := newTimeHandler("EST")
	
	//custom http server
	    server_config := &http.Server{
	    Addr: ":" + port,
		Handler: timeZone,
	    ReadTimeout: timeout ,
		WriteTimeout: timeout,
		MaxHeaderBytes: 1 << 20,
	    }
		server := server_config.ListenAndServe();
		if  server != nil{
			log.Fatal("Server failed:", server.Error())
		}
}

