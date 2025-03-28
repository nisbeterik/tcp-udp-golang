package packetHandler

import (
	"fmt"
	// "log"
	"syscall"
	"time"
)

type UDPPacket struct {
	BytesRecieved int
	ClientAddress syscall.SockaddrInet4
	Message string
	Err           error
}

func ProcessPacket(fd int, bytes int, clientAddress syscall.Sockaddr, message string, err error) {
	var packet UDPPacket
	packet.Err = err
	clientAddressIPv4, err := extractIPv4(clientAddress)
	if err != nil {
		fmt.Println("Error processing packet")
		return
	}
	packet.ClientAddress = *clientAddressIPv4
	packet.BytesRecieved = bytes
	packet.Message = message
	time.Sleep(3 * time.Second)
	printMessage(packet)
	err = sendResponse(packet.ClientAddress, fd)
	if err != nil {
		fmt.Println("Error sending response")
		return
	}
}

func extractIPv4(clientAddress syscall.Sockaddr) (*syscall.SockaddrInet4, error) {
	clientAddressIPv4, ok := clientAddress.(*syscall.SockaddrInet4)
	if !ok {
		return nil, fmt.Errorf("error extracting IPv4")
	}
	return clientAddressIPv4, nil
}

func printMessage(packet UDPPacket) {
	fmt.Printf("Client address: %d.%d.%d.%d:%d\n", packet.ClientAddress.Addr[0], packet.ClientAddress.Addr[1], packet.ClientAddress.Addr[2], packet.ClientAddress.Addr[3], packet.ClientAddress.Port)
	fmt.Printf("Bytes received: %d\n", packet.BytesRecieved)
	fmt.Println("Message from client: " + packet.Message)
}

func sendResponse(clientAddress syscall.SockaddrInet4, fd int) error {
	response := []byte("Response from server")
	err := syscall.Sendto(fd, response, 0, &clientAddress)
	if err != nil {
		return fmt.Errorf("error responding to client")
	}
	return nil
}

