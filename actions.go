package actions

import (
    "fmt"
    "regexp"
    "bufio"
)

type Config struct {
    Debug bool
    Output bufio.Writer
}


type action struct {
    Expression *Regexp
    Response string
}

type Actions struct {
    actions [][]Action
    config Config
   
func iface(list []string) []interface{} {
      vals := make([]interface{}, len(list))
      for i, v := range list { vals[i] = v }
      return vals
    }

func (a *Action) Check(i []byte) bool {
    if a.Expression.Match(0) {
        args:=a.Expression.FindSubmatch(in)
        r:=fmt.Sprintf(a.Response, iface(args))
        if Config.Debug {
            log.Printf("#Action [%s]: Response [%s]", a.Expression, r)
        }
    return true    
    }
    return false
}

func (a *Actions) CheckAll(i []byte) bool {
flag:=false
for x:=range(a.actions) {
    for y:=range(a.actions[x]) {
        if a.actions[x][y].Check(i) {
            flag=true
            if a.config.Debug {
            log.Printf("#Action priority %d position %d matched\n",y,x)
            }
        }
    }
}    
return flag
    
}
    
func Init(i bufio.Writer, debug bool) Actions {
    var a Actions
    Config.Debug=debug
    Config.Output=i
    return a
}