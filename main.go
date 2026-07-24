package main

import (
	"embed"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

//go:embed all:web
var WebFS embed.FS

//go:embed internal/icon/logo.png
var embeddedLogo []byte

var proxyServerURL string

func main() {
	runExternal := false

	dataDir, err := os.UserConfigDir()
	if err == nil {
		dataDir = filepath.Join(dataDir, "Seanime")
	}

	if os.Getenv("SEANIME_DATA_DIR") != "" {
		dataDir = os.Getenv("SEANIME_DATA_DIR")
	}

	if err == nil {
		configPath := filepath.Join(dataDir, "config.toml")

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.SetConfigFile(configPath)

		if err := viper.ReadInConfig(); err == nil {
			proxyServerURL = viper.GetString("server.proxy_server_url")
			if proxyServerURL != "" {
				runExternal = true
			}
		}
	}

	if runExternal {
		target, err := url.Parse(proxyServerURL)
		if err != nil {
			log.Fatalf("invalid target URL: %v", err)
		}

		proxy := &httputil.ReverseProxy{
			Rewrite: func(r *httputil.ProxyRequest) {
				r.SetURL(target)
				r.Out.Host = r.In.Host
			},
		}

		server := &http.Server{
			Addr:    "127.0.0.1:43211",
			Handler: proxy,
		}

		log.Printf("Proxy listening on http://127.0.0.1:43211 -> %s", target)
		log.Fatal(server.ListenAndServe())
	} else {
		server.StartServer(WebFS, embeddedLogo)
	}
}
