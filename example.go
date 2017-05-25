package main

import (

	"github.com/alknopfler/Gologger/gologger"
	"os"
	"github.com/sirupsen/logrus"
)

func main(){

	//Init with stdout and info level
	gologger.Init(os.Stdout, logrus.InfoLevel)

	//example to print an error
	gologger.Print("WARN",7,"Description to show","filename.go")



}
