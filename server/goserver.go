package goserver

import (
	"log"
	"net/http"
	"os"
	"strings"
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
	//Custom loggers for logs write to file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile| log.Lmsgprefix)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile | log.Lmsgprefix)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile | log.Lmsgprefix)
	
	//use type to bring in env vars and retrieve for use
	vars := new(dotenv)
	vars.host, vars.port = "HOST", "PORT"
	port := env.GoEnvVariables(vars.port)
	//host := env.GoEnvVariables(vars.host)
	timeout := 10 * time.Second
	//set your timezone
	timeZone := newTimeHandler("CST")
	
	//custom http server
	    server_config := &http.Server{
	    Addr: ":" + port,
		Handler: timeZone,
	    ReadTimeout: timeout ,
		WriteTimeout: timeout,
		MaxHeaderBytes: 1 << 20,
		ErrorLog: ErrorLogger,
	    }
		//check your enviroment
		var development_env string= ""
		dev := &development_env
		if strings.Contains(port, "8"){
			WarningLogger.Println("Develpment enviroment")
			*dev =  "[Env]: Development Sandbox"
		} else{
			WarningLogger.Println("Production enviroment")
			*dev = "[Env]: Production Sandbox"
		}

		//server logs
		log.Println(development_env)
		log.Println("[HTTP] Server running")
		InfoLogger.Println("[HTTP] Server running")
		//server listener
		server := server_config.ListenAndServe();
		if  server != nil{
			WarningLogger.Println("[HTTP] Server failed to run")
			log.Fatal("Server failed: ", server.Error())
		}
		return
	}