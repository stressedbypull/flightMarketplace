package config

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Postgres                      DBConfig
	ShowDocs                      bool
	Server                        Server
	Keycloak                      Keycloak
	LogHTTPRequest                bool
	DisableKeycloakAuthentication bool
}

type Server struct {
	Port string
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Schema   string
	SSLMode  string
	DBname   string
}

type Keycloak struct {
	AuthServerUrl  string
	TokenTolerance int
}

// Config contains global configuration
var Config Configuration

// KeycloakPubkey interface
var KeycloakPublicKey *rsa.PublicKey

// SetupConfig load the configuration from config.json
func SetupConfig() {
	var raw []byte
	var err error

	if raw, err = os.ReadFile("config.json"); err != nil {
		if raw, err = os.ReadFile("../config.json"); err != nil {
			log.Fatal("Unable to read configuration file: ", err)
		}
	}

	if err = json.Unmarshal(raw, &Config); err != nil {
		log.Fatal("Unable to parse configuration file: ", err)
	}
}
