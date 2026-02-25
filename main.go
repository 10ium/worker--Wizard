package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	srcPath           string
	workerPath        string
	cachePath         string
	isAndroid         = false
	VERSION           = "dev"
	defaultWorkerURLs = []string{
		"https://raw.githubusercontent.com/10ium/free-config/main/worker/iptv_player.txt",
		"https://raw.githubusercontent.com/10ium/free-config/main/worker/ClashFa_Mirror_Pro.txt",
		"https://raw.githubusercontent.com/10ium/free-config/main/worker/great_mihomo_converter_v2%2Bui.txt",
		"https://raw.githubusercontent.com/10ium/free-config/main/worker/iran_proxy.txt",
	}
)

func init() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()
	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	initPaths()
	setDNS()
	checkAndroid()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runWizard()
	}()

	server := &http.Server{Addr: ":8976"}
	http.HandleFunc("/oauth/callback", callback)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			failMessage("Error serving localhost.")
			log.Fatalln(err)
		}
	}()

	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
}
