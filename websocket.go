package main
import (
  "code.google.com/p/go.net/websocket"
  "fmt"
  "log"
  "net/http"
  "html/template"
  "io"
)

func Echo(ws *websocket.Conn) {
  var err error
  fmt.Println("WS Opening")
  for {
    fmt.Println("Listening start")
    var reply string

    if err = websocket.Message.Receive(ws, &reply); err != nil {
      if err == io.EOF {
        fmt.Println("WS Closed")
      } else {
        fmt.Println("Can't receive")
      }
      break
    }

    fmt.Println("Received back from client: " + reply)
    
    msg := "Received: " + reply
    fmt.Println("Sending to client: " + msg)
    
    if err = websocket.Message.Send(ws, msg); err != nil {
      fmt.Println("Can't send")
      break
    }
  }
}

func chat(w http.ResponseWriter, r*http.Request) {
  r.ParseForm()
  t,_ := template.ParseFiles("websocket.html")
  t.Execute(w, nil)
}

func main(){
  http.Handle("/", websocket.Handler(Echo))
  http.HandleFunc("/chat", chat)
  if err := http.ListenAndServe(":8080", nil); err != nil {
  log.Fatal("ListentAndServe:", err)
  }
}