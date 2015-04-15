package config

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	arguments := []string{"-user", "root", "-password", "toor", "-protocol", "xdcp", "-host", "lolcathost", "-port", "1337", "-dbname", "NSADB"}
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	cnf := ParseArgs(f, arguments)
	assert.Equal("root", cnf.User)
	assert.Equal("toor", cnf.Password)
	assert.Equal("xdcp", cnf.Protocol)
	assert.Equal("lolcathost", cnf.Host)
	assert.Equal(1337, cnf.Port)
	assert.Equal("NSADB", cnf.DBName)
}

func TestParseDefault(t *testing.T) {
	assert := assert.New(t)
	arguments := []string{}
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	cnf := ParseArgs(f, arguments)
	assert.Equal("tcp", cnf.Protocol)
	assert.Equal("127.0.0.1", cnf.Host)
	assert.Equal(3306, cnf.Port)
}
