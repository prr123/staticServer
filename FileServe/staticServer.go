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
//
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


/*
	programName := os.Args[0]
	errorLog := log.New(os.Stderr, "", log.LstdFlags)
	serveLog := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	addr := ":" + portStr
*/

/*
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		out := flags.Output()
		fmt.Fprintf(out, "Usage: %v [dir]\n\n", programName)
		fmt.Fprint(out, "  [dir] is optional; if not passed, '.' is used.\n\n")
		fmt.Fprint(out, "  By default, the server listens on localhost:8080. Both the\n")
		fmt.Fprint(out, "  host and the port are configurable with flags. Set the host\n")
		fmt.Fprint(out, "  to something else if you want the server to listen on a\n")
		fmt.Fprint(out, "  specific network interface. Setting the port to 0 will\n")
		fmt.Fprint(out, "  instruct the server to pick a random available port.\n\n")
		flags.PrintDefaults()
	}

	versionFlag := flags.Bool("version", false, "print version and exit")
	hostFlag := flags.String("host", "localhost", "specific host to listen on")
	portFlag := flags.String("port", "8080", "port to listen on; if 0, a random available port will be used")
	addrFlag := flags.String("addr", "localhost:8080", "full address (host:port) to listen on; don't use this if 'port' or 'host' are set")
	silentFlag := flags.Bool("silent", false, "suppress messages from output (reporting only errors)")
	corsFlag := flags.Bool("cors", false, "enable CORS by returning Access-Control-Allow-Origin header")
	tlsFlag := flags.Bool("tls", false, "enable HTTPS serving with TLS")
	certFlag := flags.String("certfile", "cert.pem", "TLS certificate file to use with -tls")
	keyFlag := flags.String("keyfile", "key.pem", "TLS key file to use with -tls")

	flags.Parse(os.Args[1:])

	if *versionFlag {
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("%v %v\n", programName, buildInfo.Main.Version)
		} else {
			errorLog.Printf("version info unavailable! run 'go version -m %v'", programName)
		}
		os.Exit(0)
	}

	if *silentFlag {
		serveLog.SetOutput(io.Discard)
	}

	if len(flags.Args()) > 1 {
		errorLog.Println("Error: too many command-line arguments")
		flags.Usage()
		os.Exit(1)
	}

	rootDir := "."
	if len(flags.Args()) == 1 {
		rootDir = flags.Args()[0]
	}

	allSetFlags := flagsSet(flags)
	if allSetFlags["addr"] && (allSetFlags["host"] || allSetFlags["port"]) {
		errorLog.Println("Error: if -addr is set, -host and -port must remain unset")
		flags.Usage()
		os.Exit(1)
	}
*/


	//need to add time outs
	addr := ":" + portStr
	srv := &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

/*
// do not understand how this function works

	// To shut the server down cleanly in tests, we register a special route
	// where we ask it to stop. A separate goroutine performs the shutdown so
	// that the server can properly answer the shutdown request without abruptly
	// closing the connection.
	shutdownCh := make(chan struct{})
	go func() {
		<-shutdownCh
		srv.Shutdown(context.Background())
	}()

	testingKey := os.Getenv("TESTING_KEY")

	mux := http.NewServeMux()
	mux.HandleFunc("/__internal/__shutdown", func(w http.ResponseWriter, r *http.Request) {
		if testingKey != "" && r.Header.Get("Static-Server-Testing-Key") == testingKey {
			w.WriteHeader(http.StatusOK)
			defer close(shutdownCh)
		} else {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
		}
	})

*/

/*
	fileHandler := serveLogger(serveLog, http.FileServer(http.Dir(rootDir)))
	if *corsFlag {
		fileHandler = enableCORS(fileHandler)
	}
*/
	mux := http.NewServeMux()
	fileHandler:= http.FileServer(http.Dir(rootDirPath))

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

/*
	scheme := "http://"
	if *tlsFlag {
		scheme = "https://"
	}

	serveLog.Printf("Serving directory %q on %v%v", rootDir, scheme, listener.Addr())

	if *tlsFlag {
		err = srv.ServeTLS(listener, *certFlag, *keyFlag)
	} else {
		err = srv.Serve(listener)
	}
*/

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

/*
// serveLogger is a logging middleware for serving. It generates logs for
// requests sent to the server.
func serveLogger(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteHost, _, _ := strings.Cut(r.RemoteAddr, ":")
		logger.Printf("%v %v %v\n", remoteHost, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// enableCORS adds a CORS response header to allow cross-origin requests.
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

*/
