package main

import "net"
import "fmt"
import "time"
import "runtime"
import "io"
import "log"
import "encoding/json"

// Not yet used
import "github.com/cloudfoundry/gosigar"

// TODO: ensure NewTicker ticks immediately, not after given Duration

func main() {
    hastur := HasturSender()

    // Send periodic updates to Hastur
    go func() {
        dur, _ := time.ParseDuration("10s")
        regChannel := time.Tick(dur)
        for _ = range regChannel {
            fmt.Println("Would register")
        }
    }()

    saddr, err := net.ResolveUDPAddr("udp", ":8125")
    if err != nil {
        panic("Couldn't resolve UDP address and port!");
    }

    ln, err := net.ListenUDP("udp", saddr)
    if err != nil {
        panic("Couldn't open UDP socket 8125!")
    }

    for {
        dec := json.NewDecoder(ln)
        for {
            var m map[string]interface{}
            if err := dec.Decode(&m); err == io.EOF {
                fmt.Println("EOF")
                break
            }
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("%v", m)

            runtime.Gosched()
        }
    }
}
