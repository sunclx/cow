package main

import (
	"github.com/BurntSushi/toml"
	"log"
	//"path/filepath"
	"time"
)

var configtoml = struct {
	RcFile      string // config file
	LogFile     string // path for log file
	AlwaysProxy bool   // whether we should alwyas use parent proxy
	LoadBalance string // select load balance mode

	TunnelAllowedPort map[string]bool // allowed ports to create tunnel

	SshServer []string

	// authenticate client
	UserPasswd     string
	UserPasswdFile string // file that contains user:passwd:[port] pairs
	AllowedClient  string
	AuthTimeout    time.Duration

	// advanced options
	DialTimeout time.Duration
	ReadTimeout time.Duration

	Core         int
	DetectSSLErr bool

	HttpErrorCode int

	dir         string // directory containing config file
	StatFile    string // Path for stat file
	BlockedFile string // blocked sites specified by user
	DirectFile  string // direct sites specified by user

	// not configurable in config file
	PrintVer        bool
	EstimateTimeout bool   // Whether to run estimateTimeout().
	EstimateTarget  string // Timeout estimate target site.

	// not config option
	saveReqLine bool // for http and cow parent, should save request line from client
}{
	RcFile:      "rc.toml", // config file
	LogFile:     "",        // path for log file
	AlwaysProxy: false,     // whether we should alwyas use parent proxy
	LoadBalance: "backup",  // select load balance mode

	TunnelAllowedPort: map[string]bool{ // allowed ports to create tunnel
		"22": true, "80": true, "443": true, // ssh, http, https
		"873": true,                                        // rsync
		"143": true, "220": true, "585": true, "993": true, // imap, imap3, imap4-ssl, imaps
		"109": true, "110": true, "473": true, "995": true, // pop2, pop3, hybrid-pop, pop3s
		"5222": true, "5269": true, // jabber-client, jabber-server
		"2401": true, "3690": true, "9418": true, // cvspserver, svn, git
	},

	SshServer: nil,

	// authenticate client
	UserPasswd:     "",
	UserPasswdFile: "", // file that contains user:passwd:[port] pairs
	AllowedClient:  "",
	AuthTimeout:    2 * time.Hour,

	// advanced options
	DialTimeout: 5 * time.Second,
	ReadTimeout: 5 * time.Second,

	Core:         1,
	DetectSSLErr: false,

	HttpErrorCode: 0,

	dir:         "/Users/chenlixin/.cow/", // directory containing config file
	StatFile:    "blocked",                // Path for stat file
	BlockedFile: "direct",                 // blocked sites specified by user
	DirectFile:  "stat",                   // direct sites specified by user

	// not configurable in config file
	PrintVer:        false,
	EstimateTimeout: false,         // Whether to run estimateTimeout().
	EstimateTarget:  "example.com", // Timeout estimate target site.

	// not config option
	saveReqLine: false, // for http and cow parent, should save request line from client

}

func main() {

	meta, err := toml.DecodeFile(configtoml.dir+configtoml.RcFile, &configtoml)
	if err != nil {
		log.Println("failed")
		log.Println(err)
		log.Println(meta)
		return
	}
	log.Println("success")
	log.Printf("congfig: %+v\n", configtoml)
	//log.Println("meta: ", meta)

}
