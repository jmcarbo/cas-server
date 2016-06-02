package main

import (
    "flag"
    "os"
    "os/signal"
    "syscall"
    "time"
    "runtime"

    "github.com/jmcarbo/cas-server/admin"
    "github.com/jmcarbo/cas-server/tools"
    "github.com/jmcarbo/cas-server/spec"
    "github.com/jmcarbo/cas-server/test"
    "github.com/jmcarbo/cas-server/storage"
)

func main() {
    // get configuration location
    var configPath string
    flag.StringVar(&configPath, "config", "", "Path to config file")
    flag.Parse()

    if configPath == "" {
        tools.Log("Command line argument `-config` must be set")
        return
    }

    config, err := tools.NewConfig(configPath)

    if err != nil {
        tools.LogError(err.Error())
        return
    }

    admin.SupportServices(config)

    storage := storage.NewMemoryStorage()

    spec.SupportV1(storage, config)

    test.SupportTest()

    startServers(config)

    // keep server running until interrupt
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
        <-c
        os.Exit(1)
    }()

    for {
        runtime.Gosched()
        time.Sleep(5 * time.Second)
    }
}
