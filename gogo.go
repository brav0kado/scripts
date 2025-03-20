// #!/usr/bin/env go run

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./gogo <input>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "process":
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./gogo process <input>")
			os.Exit(1)
		}
		input := os.Args[2]
		fmt.Println(processData(input))
	case "scan":
		scanCmd := flag.NewFlagSet("scan", flag.ExitOnError)
		portRange := scanCmd.String("ports", "1-1024", "Port range (e.g, 22,80,443 or 1-1000)")
		delayMs := scanCmd.Int("delay", 0, "Delay between scans in miliseconds")

		if len(os.Args) < 3 {
			fmt.Println("Usage: ./gogo scan <target>")
			os.Exit(1)
		}

		target := os.Args[2]
		scanCmd.Parse(os.Args[3:]) // Parse additional flags

		ports := parsePorts(*portRange)
		scanPorts(target, ports, *delayMs)

	default:
		fmt.Println("Unknown command. Available commands: process, scan")
	}
}

func processData(input string) string {
	return fmt.Sprintf("Processed: %s", input)
}

// parsePorts convers a user-defined port range into a list of integers
func parsePorts(portRange string) []int {
	var ports []int
	parts := strings.Split(portRange, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			start, err1 := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])
			if err1 == nil && err2 == nil && start <= end {
				for p := start; p <= end; p++ {
					ports = append(ports, p)
				}
			}
		} else {
			port, err := strconv.Atoi(part)
			if err == nil {
				ports = append(ports, port)
			}

		}
	}
	return ports
}

// scanPort attempts to connect to a target IP and port
func scanPort(target string, port int, delayMs int) {
	fmt.Printf("Scanning port %d on %s...\n", port, target)
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err == nil {
		fmt.Printf("[+] Port %d is OPEN\n", port)
		conn.Close()
	}
	if delayMs > 0 {
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}

// scanPorts iterates through a range of ports and scans them
func scanPorts(target string, ports []int, delayMs int) {
	fmt.Printf("Scanning ports %s...\n", target)
	for _, port := range ports {
		scanPort(target, port, delayMs)
	}
}
