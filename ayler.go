package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

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

const (
	DEBUG = true
)

// for debugging output
func linesep() {
	if DEBUG {
		fmt.Println("--------------------------")
	}
}

func dprint(a ...interface{}) {
	if DEBUG {
		fmt.Println(a...)
	}
}

const (
	state_pre_run = 0
	state_run     = 1
	state_err     = 2
	state_stopped = 3
)

// a processes managed ayler
type Process struct {
	Name    string   //Simple name
	Path    string   //path to executable
	state   int      //State of process...
	cmd     exec.Cmd //the go interface command
	io_chan chan int
}

func cTable2pTable(ctable []interface{}, ptable []Process) error {

	for i, pinfo := range ctable {
		xinfo := pinfo.(map[string]interface{})
		ptable[i].Name = xinfo["name"].(string)
		ptable[i].Path = xinfo["path"].(string)
		ptable[i].state = state_pre_run
		ptable[i].io_chan = make(chan int)
	}

	if false {
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
			if pt.Path == "" {
				break
			}

			fmt.Println(i)
			fmt.Println(pt)
		}
		linesep()
	}
	return nil
}

// run a process in a goroutine
func run_process(cmd exec.Cmd, stat_chan chan int) {
	cmd.Run()
	stat_chan <- state_stopped
}

// cTable -> unmarshalled json data
// pTable -> the actual table used by ayler to manage everything
func main() {
	// PTable => process table
	var PTable []Process
	PTable = make([]Process, 10)
	// CTable => config table
	var CTable []interface{}

	dprint("Reading configuration")
	conf, err := ioutil.ReadFile("proc.json")

	if err != nil {
		fmt.Println("\t ReadFile Error\t", err)
	}

	err = json.Unmarshal(conf, &CTable)
	if err != nil {
		fmt.Println("\t Unmarshal Error\t", err)
	}

	cTable2pTable(CTable, PTable)

	dprint("Runing processes")

	for _, pinfo := range PTable {
		if pinfo.Path != "" {
			newPath, err := exec.LookPath(pinfo.Path)
			if err != nil {
				fmt.Println("\terror with path\t", pinfo.Path)
				pinfo.state = state_err
				continue
			}
			pinfo.cmd.Path = newPath
			dprint("Running", newPath)
			go run_process(pinfo.cmd, pinfo.io_chan)
			pinfo.state = state_run
		}
	}

	// select on all the pinfo channels
	for {
		for _, pinfo := range PTable {
			if pinfo.Path != "" {
				select {
				case rv := <-pinfo.io_chan:
					pinfo.state = rv
				default:
					//nothing
				}
			}
		}
	}

}
