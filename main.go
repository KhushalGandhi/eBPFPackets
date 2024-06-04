package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// Embed the compiled eBPF object file
//
//go:embeddrop_tcp_port.o
var dropTcpPortObj []byte

func main() {
	// Allow the current process to lock memory for eBPF resources
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove memory limit: %v", err)
	}

	// Load the compiled eBPF object file
	spec, err := ebpf.LoadCollectionSpecFromReader(bytes.NewReader(dropTcpPortObj))
	if err != nil {
		log.Fatalf("Failed to load eBPF program spec: %v", err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatalf("Failed to create eBPF collection: %v", err)
	}
	defer coll.Close()

	program := coll.Programs["drop_tcp_port"]
	if program == nil {
		log.Fatalf("Failed to find eBPF program: %v", err)
	}

	// Create a BPF map to store the port number
	dropPortMap := coll.Maps["drop_port"]
	if dropPortMap == nil {
		log.Fatalf("Failed to find BPF map: %v", err)
	}

	// Set the port number to drop (default: 4040)
	port := uint16(4040)
	key := uint32(0)
	value := new(bytes.Buffer)
	if err := binary.Write(value, binary.LittleEndian, port); err != nil {
		log.Fatalf("Failed to write to buffer: %v", err)
	}
	if err := dropPortMap.Update(key, value.Bytes(), ebpf.UpdateAny); err != nil {
		log.Fatalf("Failed to update drop port map: %v", err)
	}

	// Attach the eBPF program to a network interface (e.g., eth0)
	ifaceIndex := 20 // Use the correct interface index for your setup
	xdpLink, err := link.AttachXDP(link.XDPOptions{
		Program:   program,
		Interface: ifaceIndex,
	})
	if err != nil {
		log.Fatalf("Failed to attach XDP program: %v", err)
	}
	defer xdpLink.Close()

	fmt.Printf("eBPF program attached to interface %d\n", ifaceIndex)

	// Handle signals for clean shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Exiting...")
}
