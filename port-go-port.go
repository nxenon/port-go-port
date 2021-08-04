package main

import (
	"bufio"
	"fmt"
	"github.com/akamensky/argparse"
	"net"
	"os"
	"time"
)


/*
Port-Go-Port forwards two sockets data to each other.

It's useful when :

    You can not or don't want to connect to your target directly.
    You can not connect to remote port because firewall is blocking a port on 0.0.0.0 ,but you can connect to that port with loopback IP address.

Port-Go-Port repo : https://github.com/nxenon/port-go-port
*/


var listenPort string // listening port for listening
var remotePort string // remote port for connecting
var listenIP string // listen ip for listening
var remoteIP string // remote ip for connecting
//var serviceName string // optional service name like : http

var clientSocket net.Conn = nil
var isClientSocketClosed bool = true

var remoteSocket net.Conn = nil


func main(){

	parseArgs()
	startPortGoPort()

}

func parseArgs(){

	/*
	parse script arguments
	 */

	parser := argparse.NewParser("./port-go-port", "Start Script")

	listen_port := parser.String("", "listen-port", &argparse.Options{Required: true, Help: "Listen Port Number"})

	remote_port := parser.String("", "remote-port", &argparse.Options{Required: true, Help: "Remote Port to Forward"})

	listen_ip := parser.String("", "listen-ip", &argparse.Options{Required: false, Help: "IP Address to Listen",
		Default: "0.0.0.0"})

	remote_ip := parser.String("", "remote-ip", &argparse.Options{Required: false, Help: "IP Address to Forward",
		Default: "127.0.0.1"})

	//service_name := parser.String("", "service", &argparse.Options{Required: false, Help: "[Optional] Determine Service [http]"})

	err := parser.Parse(os.Args)
	if err != nil {
		println(parser.Usage(err))
		os.Exit(0)
	}

	// set global vars
	listenPort = *listen_port
	remotePort = *remote_port
	listenIP = *listen_ip
	remoteIP = *remote_ip
	//serviceName = *service_name

	println("Listen IP : " + listenIP + " Listen Port : " + listenPort)
	println("Remote IP : " + remoteIP + " Remote Port : " + remotePort)

}

func startPortGoPort(){

	for true {
		closeBothConnections()
		// reset global vars
		clientSocket = nil
		remoteSocket = nil
		isClientSocketClosed = true

		startListening()

		// start connecting if client socket is connected
		if !isClientSocketClosed{

			connectToRemoteSocket()

		}

		time.Sleep(3 * time.Second)

	}
}

func startListening(){

	/*
	this function starts listening on a port ,for forwarding
 	*/

	listen_addr := listenIP + ":" + listenPort

	l, err := net.Listen("tcp", listen_addr)

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println("Start listening on : " + listen_addr)

	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
		}
	}(l)

	client_socket, err := l.Accept()

	if err != nil{
		println(err.Error())
		return
	}

	clientSocket = client_socket
	isClientSocketClosed = false

	println("Client Connected : " + client_socket.RemoteAddr().String())

}

func connectToRemoteSocket(){

	/*
	this function creates remote socket ,for forwarding
	 */

	remote_addr := remoteIP + ":" + remotePort

	rs, err := net.Dial("tcp", remote_addr) // rs is remote socket

	if err != nil {
		println(err.Error())
		return
	}

	remoteSocket = rs

	go forwardToRemotePort()
	forwardFromRemotePort()
}

func forwardToRemotePort() {

	/*
		this function forwards client socket data to remote socket
	*/

	for true {
		data, err := bufio.NewReader(clientSocket).ReadString('\n')
		if err != nil {
			return
		}
		processed_data := commitServiceFilters(data, "client")
		println("to remote")
		println(processed_data)
		remoteSocket.Write([]byte(processed_data))
	}

}

func forwardFromRemotePort(){

	/*
	   this function forwards remote socket data to client socket
	*/

	for true {
		data, err := bufio.NewReader(remoteSocket).ReadString('\n')
		if err != nil {
			return
		}
		processed_data := commitServiceFilters(string(data), "remote")
		println("from remote")
		println(processed_data)
		clientSocket.Write([]byte(processed_data))
	}

}

func closeBothConnections(){

	/*
		close remote and client sockets
	*/

	if clientSocket != nil {
		err := clientSocket.Close()
		if err != nil {
			println(err.Error())
		}
	}

	if remoteSocket != nil {
		err2 := remoteSocket.Close()
		if err2 != nil {
			println(err2.Error())
		}
	}

	isClientSocketClosed = true

}

func commitServiceFilters(data string, _from string)(string){

	/*
	this function process data and change them if it is necessary and if service is not None
	data is data from/to remote port before processing
	_from can be "remote" ---> from remote socket  or "client" ---> from client socket
	 */

	var processed_data string

	//if serviceName == "http"{
	//	if _from == "client"{
	//
	//		splitted := strings.Split(clientSocket.LocalAddr().String(),":")
	//
	//		replaced := strings.ReplaceAll(data, splitted[0], remoteIP)
	//		replaced2 := strings.ReplaceAll(replaced, ":" + listenPort, ":" + remotePort)
	//		processed_data = replaced2
	//
	//	} else if _from == "remote"{
	//
	//		splitted := strings.Split(clientSocket.LocalAddr().String(),":")
	//
	//		replaced := strings.ReplaceAll(data, remoteIP, splitted[0])
	//		replaced2 := strings.ReplaceAll(replaced, ":" + remotePort, ":" + listenPort)
	//		processed_data = replaced2
	//
	//	}
	//
	//} else{
	//	processed_data = data
	//}
	processed_data = data
	return processed_data

}
