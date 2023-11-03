// Server implementation.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
// https://github.com/eliben/static-server/blob/main/internal/server/server.go
//
// modifications:
// 1. replace cli
// 2. replace logger
// 3. add dbg feature
//
// v1: stripped down
// v1alt: replace httpServe with io.Copy
//

package main

import (
//	"context"
	"crypto/tls"
	"errors"
	"fmt"
//	"io"
	"log"
	"net"
	"net/http"
	"os"
//	"runtime/debug"
//	"strings"

	util "github.com/prr123/utility/utilLib"
)

type rootObj struct {
		root string
	}

func main() {


    numarg := len(os.Args)
    flags:=[]string{"dbg","port", "root"}

    useStr := "./staticServer /port=portStr /root=wwwDir [/dbg]"
    helpStr := "program that servers static files from the root directory\n"

    if numarg > len(flags)+1 {
        fmt.Println("too many arguments in cl!")
        fmt.Printf("usage: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s\n", useStr)
        os.Exit(1)
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg := false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
        fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

	portStr:=""
    portval, ok := flagMap["port"]
    if !ok {
		log.Fatalf("cli error -- port flag is required!")
    } else {
        if portval.(string) == "none" {log.Fatalf("cli error -- port value is required with /port flag!")}
       	portStr = portval.(string)
    }

	rootDirPath:=""
    rootval, ok := flagMap["root"]
    if !ok {
		log.Fatalf("cli error -- root flag is required!")
    } else {
        if rootval.(string) == "none" {log.Fatalf("cli error -- rootDir string is required with /root flag!")}
       	rootDirPath = rootval.(string)
    }

	// check whether root dir exists
	if _, err := os.Stat(rootDirPath); os.IsNotExist(err) {
		log.Fatalf("root directory: %s does not exist!", rootDirPath)
   }

	if dbg {
    	log.Printf("debug:   %t\n", dbg)
    	log.Printf("port:    %s\n", portStr)
    	log.Printf("rootDir: %s\n", rootDirPath)
	}

	fileHandler := &rootObj{
		root: rootDirPath,
	}

	//need to add time outs
	addr := ":" + portStr
	srv := &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	mux := http.NewServeMux()


	mux.Handle("/", fileHandler)
	srv.Handler = mux

	// Use an explicit listener to access .Addr() when serving on port :0
	listener, err := net.Listen("tcp", addr)
	if err != nil {
//		errorLog.Println(err)
		fmt.Printf("error -- net.Listen: %v\n", err)
		os.Exit(1)
//		return 1
	}

	err = srv.Serve(listener)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
//		errorLog.Println("Error in Serve:", err)
		fmt.Printf("error srv.Serve: %v\n", err)
		os.Exit(1)
//		return 1
	} else {
		fmt.Printf("error srv.Serve: server closed!\n")
		os.Exit(0)
//		return 0
	}

}

func (fileHandler *rootObj) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("root Dir: %s\n", (*fileHandler).root)
}
