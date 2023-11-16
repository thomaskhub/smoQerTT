package main

import (
	"crypto/tls"
	"flag"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/thomaskhub/sMoQerTT/utils"
	"go.uber.org/zap"
)

func main() {
	logger := utils.Logger{}
	logger.Init("debug")

	configFile := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	cfg := utils.ParseConfig(*configFile)

	tlsConfig := &listeners.Config{}

	if cfg.Tls.EnableTLS {
		tlsCert, err := tls.LoadX509KeyPair(cfg.Tls.TlsCrt, cfg.Tls.TlsKey)
		if err != nil {
			logger.Fatal("Failed to load cert: %v", zap.Error(err))
		}
		tlsConfig.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
		}
	}

	tls1 := listeners.NewTCP("tls1", cfg.Tls.Listening, tlsConfig)

	//Setup the websocket listener
	wsConfig := &listeners.Config{}
	if cfg.Ws.EnableTLS {
		wsCert, err := tls.LoadX509KeyPair(cfg.Ws.WsCrt, cfg.Ws.WsKey)
		if err != nil {
			logger.Fatal("Failed to load cert: %v", zap.Error(err))
		}
		wsConfig.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{wsCert},
		}
	}
	ws1 := listeners.NewWebsocket("ws1", cfg.Ws.Listening, wsConfig)

	server := mqtt.New(nil)

	server.AddListener(tls1)
	server.AddListener(ws1)

	// rule := auth.ACLRule{
	// 	Username: "websocket",
	// 	Filters: auth.Filters{
	// 		"#":        0,
	// 		"events/#": 1,
	// 	},
	// }

	server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: cfg.Ledger.Auth,
			// ACL: auth.ACLRules{
			// 	rule,
			// },
			ACL: cfg.Ledger.ACL,
		},
	})

	server.Serve()

	select {}
}
