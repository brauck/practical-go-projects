// basic-http-server/serve-html/main.go
package main

import (
	"log"
	"log/syslog"
	"os"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	log.Println("System is starting")
	log.Println("Retrieving user 12")
	log.Println("Compute the total of invoice 1262663663")

	logwriter, err := syslog.New(syslog.LOG_WARNING|syslog.LOG_DAEMON, "loggingTestProgram")
	if err != nil {
		log.Fatal(err)
	}
	logwriter.Emerg("emergency sent to syslog. TEST2")
}
