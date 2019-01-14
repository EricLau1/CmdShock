package utils

import(
  "io/ioutil"
)

func Write(filename, content string) error {

  return ioutil.WriteFile(filename, []byte(content), 0664)

}
