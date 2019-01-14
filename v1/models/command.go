package models

import(
  "fmt"
  "errors"
)

var(
  ErrCommandNameTaken = errors.New("Comando ja existe")
  ErrCommandEmptyField = errors.New("Campos Vazios")
)

type Command struct {
  Id int
  Name string
  Description string
  Terminal Terminal
  CreatedAt string
}

func isValid(cmd Command) (bool, error) {
  
  if IsEmpty(cmd) {
  
     return false, ErrCommandEmptyField 

  }

  count, err := Unique("commands", "name", cmd.Name)

  if err != nil {
  
    return false, err

  }
  
  if count > 0 {
    
    return false, ErrCommandNameTaken

  }

  return true, nil

}

func NewCommand(cmd Command) (bool, error) {

  validate, err := isValid(cmd)

  if err != nil {
    
    return validate, err

  }

  con := Connect()

  sql := "insert into commands (name, description, terminal) values (?, ?, ?)"

  stmt, err := con.Prepare(sql)

  if err != nil {
  
    return false, err

  }

  _, err = stmt.Exec(cmd.Name, cmd.Description, cmd.Terminal.Id)

  if err != nil {
  
    return false, err

  }

  defer stmt.Close()
  defer con.Close()

  return true, nil
}

func GetCommands()([]Command, error) {

  con := Connect()

  sql := "select * from commands"

  rs, err := con.Query(sql)

  if err != nil {
  
    return nil, err

  }
  
  var commands []Command
  
  for rs.Next() {
    
    var command Command

    var terminalId int

    err := rs.Scan(&command.Id, &command.Name, &command.Description, &terminalId, &command.CreatedAt)

    if err != nil {
  
      return nil, err

    }

    terminal, err := GetTerminal(terminalId)
    
    if err != nil {
  
      return nil, err

    }

    command.Terminal = terminal

    commands = append(commands, command)

  }

  defer rs.Close()
  defer con.Close()

  return commands, nil
}

func SearchCommands(search string)([]Command, error){

  con := Connect()

  search = fmt.Sprintf("%%%s%%", search);
  
  sql := "select * from commands where name like ?"

  rs, err := con.Query(sql, search)

  if err != nil {
    
    return nil, err

  }

  var commands []Command
  
  for rs.Next() {
    
    var command Command
    var terminalId int

    err := rs.Scan(&command.Id, &command.Name, &command.Description, &terminalId, &command.CreatedAt)
    
    if err != nil {
    
      return nil, err

    }

    command.Terminal, err = GetTerminal(terminalId)

    if err != nil {
    
      return nil, err

    }
    
    commands = append(commands, command)
  }

  defer rs.Close()
  defer con.Close()

  return commands, nil
}
