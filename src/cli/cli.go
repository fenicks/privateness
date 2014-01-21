// Command line interface arguments
package cli

import (
    "flag"
    "github.com/op/go-logging"
    "github.com/skycoin/skycoin/src/util"
    "log"
)

type Args interface {
    register()
    postProcess()
    getConfig() *Config
}

type Config struct {
    DisableGUI   bool
    DisableCoind bool
    // DHT uses this port for UDP; gnet uses this for TCP incoming and outgoing
    Port int
    // Remote web interface
    EnableWebInterface bool
    WebInterfacePort   int
    // Data directory holds app data -- defaults to ~/.skycoin
    DataDirectory string
    // Logging
    LogLevel logging.Level
    ColorLog bool
    // This is the value registered with flag, it is converted to LogLevel
    // after parsing
    logLevel string

    /* Developer options */

    // Enable cpu profiling
    ProfileCPU bool
    // Where the file is written to
    ProfileCPUFile string
    // HTTP profiling interface (see http://golang.org/pkg/net/http/pprof/)
    HTTPProf bool
    // Will force it to connect to this ip:port, instead of waiting for it
    // to show up as a peer
    ConnectTo string
}

func (self *Config) register() {
    log.Panic("Config.register must be overridden")
}

func (self *Config) postProcess() {
    self.DataDirectory = util.InitDataDir(self.DataDirectory)
    ll, err := logging.LogLevel(self.logLevel)
    if err != nil {
        log.Panic("Invalid -log-level %s: %v\n", self.logLevel, err)
    }
    self.LogLevel = ll
}

func (self *Config) getConfig() *Config {
    return self
}

// Parses arguments defined in a struct that satisfies Config interface
func ParseArgs(args Args) *Config {
    args.register()
    flag.Parse()
    args.postProcess()
    return args.getConfig()
}
