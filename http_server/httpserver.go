package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

type LoggerServer struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (l *LoggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(l.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(l.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(l.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(l.LogWriter, "-----------------------\n")
	l.Handler.ServeHTTP(w, r)
}

type BasicAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (b *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if ok {
		if user == b.User && pass == b.Password {
			b.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintf(w, "User or password incorrect\n")
		}
	} else {
		fmt.Fprintf(w, "Error trying to retrieve data from Basic auth")
	}
}

func main() {
	fmt.Println("Enter the type of server you want to launch from following:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and authentication")
	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	var mySuperServer http.Handler
	switch selection {
	case 1:
		mySuperServer = new(MyServer)
	case 2:
		mySuperServer = &LoggerServer{
			Handler:   new(MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string
		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)

		mySuperServer = &LoggerServer{
			Handler: &BasicAuthMiddleware{
				Handler:  new(MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(MyServer)
	}
	http.Handle("/", mySuperServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
