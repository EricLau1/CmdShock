package routes

import (
  "net/http"
  "github.com/gorilla/mux"
  "log"
  "fmt"
  "strconv"
  "strings"
  "../utils"
  "../models"
)

func NewRouter() *mux.Router {

  r := mux.NewRouter()

  r.HandleFunc("/", homeGetHandler).Methods("GET")
  r.HandleFunc("/commands", commandsGetHandler).Methods("GET")
  r.HandleFunc("/commands", commandsPostHandler).Methods("POST")
  r.HandleFunc("/search", commandsSearchGetHandler).Methods("GET")
  r.HandleFunc("/dump", dumpGetHandler).Methods("GET")

  fileServer := http.FileServer(http.Dir("./assets"))

  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

  return r

}

func homeGetHandler(w http.ResponseWriter, r *http.Request){
  
  utils.ExecuteTemplate(w, "index.html", nil)

}

func commandsGetHandler(w http.ResponseWriter, r *http.Request) {
  
  commands, err := models.GetCommands()

  if err != nil {

    log.Fatal(err.Error())
    utils.InternalServerError(w)
    return

  }
  
  terminals, err := models.GetTerminals()
  
  if err != nil {

    log.Fatal(err.Error())
    utils.InternalServerError(w)
    return
  }

  var displayTable bool = false 

  if len(commands) > 0 {
    displayTable = true
  }

  utils.ExecuteTemplate(w, "command_list.html", struct{ 
    Terminals []models.Terminal
    Commands []models.Command
    DisplayTable bool
  }{
    Terminals: terminals,
    Commands: commands,
    DisplayTable: displayTable,
  })
}

func commandsPostHandler(w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  var command models.Command
  command.Name = r.PostForm.Get("name")
  command.Description = r.PostForm.Get("description")
  terminalId, _ := strconv.Atoi(r.PostForm.Get("terminal")) 
  terminal, err := models.GetTerminal(terminalId)
  
  if err != nil {
    utils.InternalServerError(w)
    return
  }

  command.Terminal = terminal

  _, err = models.NewCommand(command)

  if err != nil {
    
    switch(err) {
      case models.ErrCommandNameTaken:
        fmt.Println("comando ja existe!")
        break;
      case models.ErrCommandEmptyField:
          fmt.Println("campos em branco!")
      default:
        utils.InternalServerError(w)
      return
    }
  }

  http.Redirect(w, r, "/commands", 302)
}

func commandsSearchGetHandler(w http.ResponseWriter, r *http.Request) {
  
  keys := r.URL.Query()
  
  search := keys.Get("command")

  if search == "" {
  
    http.Redirect(w, r, "/commands", 302)
    return

  }

  commands, err := models.SearchCommands(search)

  if err != nil {
  
    utils.InternalServerError(w)
    return

  }

  var results int = len(commands)

  utils.ExecuteTemplate(w, "command_search.html", struct{
    Search string
    Results int
    Commands []models.Command
  }{
    Search: search,
    Results: results,
    Commands: commands,
  })  
}

func dumpGetHandler(w http.ResponseWriter, r *http.Request) {
  
  commands, err := models.GetCommands()

  if err != nil {
    utils.InternalServerError(w)
    return
  }

  var sql string = ""

  sql += "insert into os (platform, architecture) values ('Linux', 'x64');\r\n"
  sql += "insert into terminal (name, os) values ('bash', 1);\r\n"
  sql += "insert into commands (name, description, terminal) values \r\n"
  for _, command := range commands {

    command.Name = strings.Replace(command.Name, "'", "\"", -1)
    command.Description = strings.Replace(command.Description, "'", "\"", -1)

    sql += fmt.Sprintf("('%s','%s',%d),\r\n", command.Name, command.Description, command.Terminal.Id)
  
  }

  sql = sql[:len(sql) - 3]
  sql += ";"

  fmt.Println(sql)

  err = utils.Write("dump.sql", sql)

  if err != nil {
  
    fmt.Println(err)
    w.Write([]byte("dump error..."))
  
  }
  w.Write([]byte("dump complete!"))
}
