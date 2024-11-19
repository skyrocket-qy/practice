package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"sync"
)

// Struct to represent a peer
type Peer struct {
	Addr string
	Port string
}

// Start the peer server to handle file-sharing requests
func startServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	ln, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	log.Printf("Server listening on port %s...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

// Handle incoming connection and file transfer request
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Receive the file request (name of the file)
	var fileName string
	_, err := fmt.Fscan(conn, &fileName)
	if err != nil {
		log.Println("Error reading file name:", err)
		return
	}

	// Check if the requested file exists
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			conn.Write([]byte("File not found"))
			return
		}
		log.Println("Error checking file:", err)
		return
	}

	// Open the file and send it to the requesting peer
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the file content to the client
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file:", err)
	}
	log.Printf("Sent file: %s\n", fileName)
}

// Connect to a peer and request a file
func requestFile(peerAddr, fileName string) error {
	conn, err := net.Dial("tcp", peerAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the file request to the peer
	fmt.Fprintln(conn, fileName)

	fileContent, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fileContent.Close()

	_, err = io.Copy(fileContent, conn)
	if err != nil {
		return err
	}

	log.Printf("File downloaded: %s\n", fileName)
	return nil
}

// Function to make the client both a server and client
func startClientMode(serverPort string, wg *sync.WaitGroup) {
	wg.Add(1)

	// Start the server mode (peer server listening for file requests)
	go startServer(serverPort, wg)
}

func main() {
	// Example peer server and peer client setup
	peerAddr := "localhost:8000" // Address of another peer to request files from
	serverPort := "8001"         // Port to listen on for incoming requests

	var wg sync.WaitGroup
	startClientMode(serverPort, &wg)
	exp := regexp.MustCompile("^get .+")
	var action string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please enter the Action: ")
		scanner.Scan()
		action = scanner.Text()
		if exp.MatchString(action) {
			fileName := action[4:]

			if fileName == "" {
				fmt.Println("Filename cannot be empty")
			} else {
				fmt.Printf("Requested filename: %s\n", fileName)
				requestFile(peerAddr, fileName)
			}
		} else {
			fmt.Println("Invalid action. Please enter a valid action like 'get <filename>'")
		}

	}
}
