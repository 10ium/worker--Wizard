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

type WorkerSource struct {
	Name       string
	URL        string
	IsLegacy   bool
	IsCustom   bool
	KVBindings []string
	VarPrompts []EnvVarPrompt
}

type EnvVarPrompt struct {
	Name     string
	Prompt   string
	IsSecret bool
}

var (
	srcPath       string
	workerPath    string
	cachePath     string
	isAndroid     = false
	VERSION       = "dev"
	workerSources = []WorkerSource{
		{
			Name:       "Original BPB Worker (legacy mode)",
			URL:        "https://github.com/bia-pain-bache/BPB-Worker-Panel/releases/latest/download/worker.js",
			IsLegacy:   true,
			KVBindings: []string{"kv"},
		},
		{
			Name: "iptv_player",
			URL:  "https://raw.githubusercontent.com/10ium/free-config/main/worker/iptv_player.txt",
		},
		{
			Name: "ClashFa_Mirror_Pro",
			URL:  "https://raw.githubusercontent.com/10ium/free-config/main/worker/ClashFa_Mirror_Pro.txt",
		},
		{
			Name: "great_mihomo_converter_v2+ui",
			URL:  "https://raw.githubusercontent.com/10ium/free-config/refs/heads/main/worker/great_mihomo_converter",
		},
		{
			Name: "iran_proxy",
			URL:  "https://raw.githubusercontent.com/10ium/free-config/main/worker/iran_proxy.txt",
		},
		{
			Name:       "TG-Proxy_Assistant",
			URL:        "https://raw.githubusercontent.com/10ium/free-config/refs/heads/main/worker/TG-Proxy_Assistant",
			KVBindings: []string{"PROXY_KV"},
			VarPrompts: []EnvVarPrompt{
				{Name: "BOT_TOKEN", Prompt: "- Please enter your Telegram bot token (created by BotFather): ", IsSecret: true},
				{Name: "ADMIN_PASSWORD", Prompt: "- Please set an admin password: ", IsSecret: true},
			},
		},
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
