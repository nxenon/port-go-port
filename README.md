# Port-Go-Port
[![Tool Category](https://badgen.net/badge/Tool/Socket%20Forwarder/black)](https://github.com/nxenon/port-go-port)
[![APP Version](https://badgen.net/badge/Version/Beta/red)](https://github.com/nxenon/port-go-port)
[![Go Version](https://badgen.net/badge/Go/1.13/blue)](https://golang.org/doc/go1.13)
[![License](https://badgen.net/badge/License/GPLv2/purple)](https://github.com/nxenon/port-go-port/blob/master/LICENSE)

**Port-Go-Port** forwards two sockets data to each other.

It's useful when :
- You can not or don't want to connect to your target directly.
- You can not connect to remote port because firewall is blocking a port on 0.0.0.0 ,but you can connect to that port with `loopback` IP address.

Installation & Build
----
    git clone https://github.com/nxenon/port-go-port.git
    cd port-go-port
    Then you should compile the port-go-port.go file:
    for windows run:
        GOOS=windows go build port-go-port.go
    for linux run:
        GOOS=linux go build port-go-port.go

    Then you can run the executable.

Usage
---
    ./port-go-port --listen-ip 0.0.0.0 --listen-port 5678 --remote-ip 192.168.1.1 --remote-port 23
    then connect to port 5678 to start forwarding

Help
----

    usage: ./port-go-port [-h|--help] [--listen-port] [--remote-port] [--listen-ip] [--remote-ip]
    
            Start The Script
    
    Arguments:
    -h  --help         Print help information
    --listen-port  Listen Port Number
    --remote-port  Remote Port to Forward
    --listen-ip    IP Address to Listen. Default: 0.0.0.0
    --remote-ip    IP Address to Forward. Default: 127.0.0.1

