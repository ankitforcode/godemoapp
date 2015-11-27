package main

import (
  "github.com/Sirupsen/logrus"
  logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
  "io"
  "log"
  "log/syslog"
  "net/http"
  "os"
)

// Function Home - that handles a web server request for a page
func Home(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello World!")
  log := logrus.New()
  hook, err := logrus_syslog.NewSyslogHook("udp", os.Getenv("SYSLOG_SERVER")+":"+os.Getenv("SYSLOG_PORT"), syslog.LOG_INFO, "")
  if err == nil {
    log.Hooks.Add(hook)
  }
  log.Info(req, "\n")
}

// Function: MAIN
func main() {
  // start up output - to let us know the web server is starting
  log.Println("Starting server...")
  // handler registered for requests to the root of our web server - send them to the function 'Home'
  http.HandleFunc("/", Home)
  // start a web server on localhost port 8000 grabbing any errors that occur
  err := http.ListenAndServe(":8000", nil)
  // check for any errors from the above web server start-up
  if err != nil {
    log.Fatal("ListenAndServer: ", err)
  }
}
