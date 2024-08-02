package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"node-x/utils"
	"os"
	"time"

	"node-x/asset"

	"gitee.com/dark.H/gs"
)

var explorer *utils.Console
var err error
var RUNNING_AT = time.Now().Format("2006-01-02 15:04:05")
var PORT = 31111

func main() {
	updater := ""
	daemon := false
	flag.StringVar(&updater, "update_me_by_updater_this_is_a_tag_has_no_meaning", "", "updater path")
	flag.BoolVar(&daemon, "d", false, "true to run daemon")
	flag.Parse()

	if daemon {
		utils.DaemonLog([]string{os.Args[0]}, "/tmp/node-x.log")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}
	explorer, err = utils.NewBrowserContext()
	if err != nil {
		fmt.Println("Failed to create browser context:", err)
		return
	}
	Release()
	LoadALlConfig()
	RUNNING_AT = time.Now().Format("2006-01-02 15:04:05")

	http.HandleFunc("/v1/web", webHandler)
	http.HandleFunc("/v1/text", webTextHandler)
	http.HandleFunc("/v1/channel", webChannelHandler)
	http.HandleFunc("/v1/link", weblinkHandler)
	http.HandleFunc("/v1/link/stream", weblinkStreamHandler)
	http.HandleFunc("/v1/news", webNewsHandler)
	http.HandleFunc("/v1/news/stream", webNewsStreamHandler)
	http.HandleFunc("/v1/burp", webBurpHandler)
	http.HandleFunc("/v1/raw", rawHandler)
	http.HandleFunc("/v1/info", infoHandler)
	http.HandleFunc("/v1/join", reciveconnect)
	http.HandleFunc("/v1/install", installHandler)
	http.HandleFunc("/v1/update", updateInstaller)
	http.HandleFunc("/v1/upgrade", upgradeHandler)
	http.HandleFunc("/v1/config", configupdateInstaller)
	http.HandleFunc("/v1/png", downloadImgHandler)
	http.HandleFunc("/v1/log", logHandler)
	http.HandleFunc("/v1/search", searcherHandler)
	http.HandleFunc("/v1/config/upload", jsUpdateInstaller)
	time.Sleep(3 * time.Second)
	// fmt.Println("Server started on :31111")
	RunServer(PORT)
	// http.ListenAndServe(":31111", nil)
}

func RunServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	srv := &http.Server{Addr: addr, Handler: http.DefaultServeMux}
	// return server.ListenAndServeTLS(certFile, keyFile)
	certPEMBlock, err := asset.Asset("Res/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	keyPEMBlock, err := asset.Asset("Res/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	// srv.TLSConfig.Certificates =
	gs.Str(":").Add(port).Color("g").Color("B").Println("Server")
	for {

		func() {
			l, err := net.Listen("tcp", addr)
			if err != nil {
				fmt.Println("Err Wait 5sec")
				return
			}
			defer l.Close()
			tlsListener := tls.NewListener(l, config)
			defer tlsListener.Close()
			err = srv.Serve(tlsListener)
			if err != nil {
				fmt.Println("Err Wait 5sec")
				time.Sleep(5 * time.Second)
				err = nil
			}
		}()
		time.Sleep(5 * time.Second)
		continue
	}
}
