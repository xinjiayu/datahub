package main

import (
	"fmt"
	"github.com/asiainfoLDP/datahub/client"
	"github.com/asiainfoLDP/datahub/daemon"
	"github.com/asiainfoLDP/datahub/daemon/daemonigo"
	flag "github.com/asiainfoLDP/datahub/utils/mflag"
	"os"
)

var (
	runDaemon bool
)

func init() {
	flagParse()
}

func flagParse() {
	flDaemon := flag.Bool([]string{"D", "-daemon"}, false, "Enable daemon mode")
	flVersion := flag.Bool([]string{"V", "-version"}, false, "Show version")
	flToken := flag.String([]string{"-token"}, "", "user token")

	flag.Usage = client.ShowUsage
	//flag.PrintDefaults()
	flag.Parse()
	//fmt.Printf("run daemon: %v, version: %v\n", *flDaemon, *flVersion)

	if *flVersion {
		fmt.Println("datahub v0.4")
		os.Exit(0)
	}

	fmt.Println("token:", *flToken)

	if len(*flToken) == 40 {
		daemonigo.Token = *flToken
		daemon.DaemonID = *flToken
	}

	if *flDaemon {
		runDaemon = true

	}
}

func main() {

	if runDaemon {
		daemon.RunDaemon()
	} else {
		client.RunClient()
	}
}
