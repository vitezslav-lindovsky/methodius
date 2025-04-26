package main

import (
	"fmt"
	"github.com/spf13/cobra"
	methodius "github.com/vitezslav-lindovsky/methodius/internal"
	"log"
	"net/http"
)

const Description = `
Methodius | https://vitezslav-lindovsky.cz
------------------------------------------------------------
They: "Please describe functionality of HTTP methods".
Me:   "...but do you know what QUIT does?"
------------------------------------------------------------
--rfc for follow RFC-7231 | -v --verbose | -p --port [port]
`

const (
	RFC7231     = "RFC-7231" // https://datatracker.ietf.org/doc/html/rfc7231#section-4.3
	DefaultPort = 8080
)

func main() {
	fmt.Print(Description)
	rfc, verbose, port, err := parseFlags()

	if err != nil {
		log.Fatal(err)
	}

	bindTo := fmt.Sprintf("localhost:%d", port)
	methodsToActions := methodius.GetMethodMaps(rfc)
	methodius.PrintUsage(methodsToActions, port)
	store := methodius.NewKeyValueStore()
	server := methodius.NewServer(store, verbose, methodsToActions)
	http.HandleFunc("/", server.HandleRequest)
	fmt.Printf("Server starting on localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(bindTo, nil))
}

func parseFlags() (rfc bool, verbose bool, port int, err error) {
	var rootCmd = &cobra.Command{}

	rootCmd.Flags().BoolVar(&rfc, "rfc", false, "Follow "+RFC7231)
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose")
	rootCmd.Flags().IntVarP(&port, "port", "p", DefaultPort, "Local HTTP port")

	err = rootCmd.Execute()

	return
}
