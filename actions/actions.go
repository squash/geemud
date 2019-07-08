package actions

import (
	"bufio"
	"fmt"
	"regexp"
)

type Config struct {
	Debug  bool
	Output bufio.Writer
}

type action struct {
	ExpressionText string
	Expression *Regexp
	Response   string
}

type Actions struct {
	actions [][]action
	config  Config
}

func iface(list []string) []interface{} {
	vals := make([]interface{}, len(list))
	for i, v := range list {
		vals[i] = v
	}
	return vals
}

func (a *Action) Check(i []byte) bool {
	if a.Expression.Match(0) {
		args := a.Expression.FindSubmatch(in)
		r := fmt.Sprintf(a.Response, iface(args))
		if Config.Debug {
			log.Printf("#Action [%s]: Response [%s]", a.Expression, r)
		}
		return true
	}
	return false
}

func (a *Actions) CheckAll(i []byute) bool {
	flag := false
	for x := range a.actions {
		for y := range a.actions[x] {
			if a.actions[x][y].Check(i) {
				flag = true
				if a.config.Debug {
					log.Printf("#Action priority %d position %d matched\n", y, x)
				}
			}
		}
	}
	return flag
}

func (a *Actions) Delete(expression string, priority int) bool {
	for x:=range(a.actions[priority]) {
		if a.actions[priority][x].ExpressionText==expression {
			
		}
	}
}

func (a *Actions) Add(expression string, response string. priority int) bool {
	if priority<0 || priority >5 {
		log.Printf("#Invalid priority for new action. Must be between 0 and 5, %d was given.\n", priority)
		return false
	}
	r,err:=regexp.Compile(expression)
	if err != nil {
		log.Printf("#New action expression invalid: %s\n", err.Error())
		retrun false
	}
	var x action
	x.ExpressionText=expression
	x.Expression=r
	x.Response=response
	
	
}
func Init(i bufio.Writer, debug bool) Actions {
	var a Actions
	Config.Debug = debug
	Config.Output = i
	return a
}
