package main

import (
	log "../log"

	"fmt"
	"os"
)

func clog(){
	err := log.New(log.CONSOLE, log.ConsoleConfig{})
	// open a logger writer of console or file mode.
	/*mode := "console"
	config := `{"level":1,"filename":"test.log"}`
	log.NewLogger(0, mode, config,true)*/
	if err != nil {
		fmt.Printf("Fail to create new logger: %v\n", err)
		os.Exit(1)
	}
	//log.SetShowLine(true)
	log.Trace("-1-Hello %s!", "Clog")
	// Output: Hello Clog!

	log.Info("-2-Hello %s!", "Clog")
	fmt.Print("--------before log error skip 0-----------\n")
	log.Error(0,"-3-Hello %s!", "Clog")
	fmt.Print("--------before log error skip 1-----------\n")
	log.Error(1,"-5-Hello %s!", "Clog")
	fmt.Print("--------before log error skip 2-----------\n")
	log.Error(2,"-4-Hello %s!", "Clog")
	fmt.Print("-------- log error testtest-----------\n")
	log.Error(2,"-6-Hello %s!", "Clog")



}
func main() {

	clog()
}

