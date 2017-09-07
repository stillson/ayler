package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// use json for config

/*
null
bool
float
string
array
k/v (map)(string key, any val (islk))

process is a k/v

{
"name":"web",
"path":"/usr/local/bin/nginx"
}
*/

const DEBUG = true

// for debugging output
func linesep() {
	if DEBUG {
		fmt.Println("--------------------------")
	}
}

const (
	state_pre_run = 0
	//state_run     = 1
	//state_err     = 2
	//state_stopped = 3
)

// a processes managed ayler
type Process struct {
	Name  string //Simple name
	Path  string //path to executable
	state int    //State of process...
}

func cTable2pTable(ctable []interface{}, ptable []Process) error {

	for i, pinfo := range ctable {
		var xinfo = pinfo.(map[string]interface{})

		ptable[i].Name = xinfo["name"].(string)
		ptable[i].Path = xinfo["path"].(string)
		ptable[i].state = state_pre_run
	}

	if DEBUG {
		linesep()
		fmt.Println("ctable\n")
		linesep()
		for i, pinfo := range ctable {
			fmt.Println(i)
			fmt.Println(pinfo)
			fmt.Printf("pinfo: %T\n", pinfo)
			var xinfo = pinfo.(map[string]interface{})
			fmt.Printf("xinfo: %T\n", xinfo["name"])
		}

		linesep()
		fmt.Println("ptable")
		linesep()
		for i, pt := range ptable {
			fmt.Println(i)
			fmt.Println(pt)
		}
		linesep()
	}

	return nil
}

// cTable -> unmarshalled json data
// pTable -> the actual table used by ayler to manage everything
func main() {
	var PTable []Process
	PTable = make([]Process, 10)
	var CTable []interface{}

	var err error
	var conf []byte

	fmt.Println("Reding configuration")
	conf, err = ioutil.ReadFile("proc.json")

	if err != nil {
		fmt.Println("\t ReadFile Error")
		fmt.Println("\t", err)
	}

	err = json.Unmarshal(conf, &CTable)
	if err != nil {
		fmt.Println("\t Unmarshal Error")
		fmt.Println("\t", err)
	}

	linesep()
	cTable2pTable(CTable, PTable)
	linesep()

	fmt.Println("Runing processes")

}
