package config

import (
	"flag"
	"os"
)

const (
	defaultProtocol = "tcp"
	defaultHost     = "127.0.0.1"
	defaultPort     = 3306
)

type Config struct {
	User     string
	Password string
	Protocol string
	Host     string
	Port     int
	DBName   string
}

func Parse() Config {
	return parse(flag.CommandLine, os.Args[1:])
}

func parse(f *flag.FlagSet, arguments []string) Config {
	cnf := Config{}
	f.StringVar(&cnf.User, "user", "", "DB connection's username")
	f.StringVar(&cnf.Password, "password", "", "DB connection's password")
	f.StringVar(&cnf.Protocol, "protocol", defaultProtocol, "DB connection's protocol (tcp, unix)")
	f.StringVar(&cnf.Host, "host", defaultHost, "DB connection's host")
	f.IntVar(&cnf.Port, "port", defaultPort, "DB connection's port")
	f.StringVar(&cnf.DBName, "dbname", "", "DB name")
	f.Parse(arguments)

	return cnf
}
