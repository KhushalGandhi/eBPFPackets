# eBPF Program to Drop TCP Packets

This project contains an eBPF program written in C that drops TCP packets on a specified port. The port number is configurable from user space. Additionally, a Go program is included to interact with the eBPF program.

## Prerequisites

To run this project, you will need:

- **Docker:** A platform to develop, ship, and run applications in containers.
- **Go:** A statically typed, compiled programming language designed for building scalable and reliable software (version 1.18 or later).
- **Clang:** A compiler for the C language family that can compile eBPF programs.
- **Linux headers:** Required to compile and link eBPF programs against the Linux kernel.

## Setup

### Step 1: Clone the Repository

Clone this repository to your local machine:

```sh
git clone https://github.com/your-repo/ebpf-drop-tcp.git
cd ebpf-drop-tcp
