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
cd ebpf-drop-tcp  ```sh



Problem Statement 3: Explain the code snippet

Explain what the following code is attempting to do?
The code is creating a buffered channel cnp that can hold functions. It then starts four goroutines, each of which will execute any function received from the channel cnp. The main function then sends a function to the channel cnp and prints "Hello".

Explaining how the highlighted constructs work:

make(chan func(), 10): This creates a buffered channel that can hold up to 10 function values. The channel is used to pass functions between goroutines.
for i := 0; i < 4; i++ { go func() { for f := range cnp { f() } }() }: This loop starts four goroutines. Each goroutine waits for functions to be sent to the channel cnp and executes them.
Giving use-cases of what these constructs could be used for:

Buffered channels can be used to manage a pool of tasks that need to be processed concurrently. In this case, functions are tasks that the goroutines will execute.
Goroutines can be used to perform concurrent operations, improving the efficiency of CPU-bound or I/O-bound tasks.
What is the significance of the for loop with 4 iterations?
The loop with 4 iterations starts four separate goroutines, allowing up to four functions to be processed concurrently from the channel cnp.

What is the significance of make(chan func(), 10)?
This creates a buffered channel with a capacity of 10. It allows up to 10 functions to be queued up for execution without blocking the main function. The buffer ensures that the main function can send functions to the channel even if the goroutines are busy.

Why is “HERE1” not getting printed?
The function containing the fmt.Println("HERE1") statement is sent to the channel cnp, but the program terminates before any of the goroutines have a chance to execute the function. Since the main function does not wait for the goroutines to complete, the program ends, and "HERE1" is not printed.

