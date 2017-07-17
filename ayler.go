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

etc

*/

func linesep() {
	fmt.Println("--------------------------")
}

type Process struct {
	Name  string
	Path  string
	state int
}

func cTable2pTable(ctable []interface{}, ptable []Process) error {

	//var i int
	//var pinfo map[string]string

	for i, pinfo := range ctable {
		fmt.Println(i)
		fmt.Println(pinfo)
	}

	for i, pt := range ptable {
		fmt.Println(i)
		fmt.Println(pt)
	}

	return nil
}

func main() {
	var PTable []Process
	PTable = make([]Process, 10)
	var CTable []interface{}
	//var CTable []map[string]string

	var err error
	var conf []byte

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

	fmt.Println(PTable)
	fmt.Println("")
	fmt.Println(PTable[0])
	fmt.Println("")
	fmt.Println(CTable)
	fmt.Println("")
	fmt.Printf("%T\n", CTable)
}
