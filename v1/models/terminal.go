package models

type Terminal struct {
  Id int
  Name string
  Os int
}

func GetTerminals()([]Terminal, error) {

  con := Connect()

  sql := "select * from terminal"

  rs, err := con.Query(sql)

  if err != nil {
  
    return nil, err

  }

  var terminals []Terminal

  if rs.Next() {
    var terminal Terminal

    err := rs.Scan(&terminal.Id, &terminal.Name, &terminal.Os)
       
    if err != nil {
  
      return nil, err

    }

    terminals = append(terminals, terminal)
    
  }

  defer rs.Close()
  defer con.Close()

  return terminals, nil

}

func GetTerminal(id int)(Terminal, error) {

  con := Connect()

  sql := "select * from terminal where id = ?"

  rs, err := con.Query(sql, id)

  if err != nil {
  
    return Terminal{}, err
  }
  
  var terminal Terminal

  if rs.Next() {
  
    err := rs.Scan(&terminal.Id, &terminal.Name, &terminal.Os)

    if err != nil {
  
      return Terminal{}, err
    }

  }

  defer rs.Close()
  defer con.Close()

  return terminal, nil

}
