package main

import (
    "fmt"
    "github.com/Pluto-tv/export-parameters-vault/pkg/config"
    "github.com/Pluto-tv/export-parameters-vault/pkg/ssm"
)

func main() {
    fmt.Println("Initializing...")
    ssm.InitSsm("preprod")
    config.InitConfig("http://localhost:5071", "this-is-a-local-token-for-client-authorization")

    fmt.Println("Getting SSM keys...")
    keys, err := ssm.ListAllParameters()
    if err != nil {
        panic(err)
        return
    }

    fmt.Println("Get keys now reading values of each key...")
    for _, key := range keys {
        val, err := ssm.ReadParameterValue(key)
        if err != nil {
            panic(err)
            return
        }

        err = config.Insert(key, val)
        if err != nil {
            panic(err)
            return
        }
        fmt.Print(".")
    }

    fmt.Println("\nComplete!")
}
