# websshket

Inspired by [google/huproxy](https://github.com/google/huproxy), this aims to behave similarly, with some differences.

## Goals
Much like huproxy, we want an ssh client config which contains the following definition to tunnel our ssh connection over a websocket connection to our server.
```
Host shell.example.com
    ProxyCommand /path/to/websshket -auth=token wss://proxy.example.com/proxy/%h/%p
```

### Small steps
A small number of steps that will form the Proof-of-Concept
- Have the command line flag switch the operation mode from client to server
- Have a server that listens on a websocket and prints all binary data it receives to stdout
- Make the server initiate a tcp connection to the upstream based on the handled url, e.g. `/proxy/%h/%p`
- Make the server pass all binary data over to the established tcp connection instead of printing to stdout
- Make the server read all response from the tcp connection and write to the websocket connection
- Build the client that can connect and read/write to the websocket connection

There are further goals, which involve developing the authentication token system, and adding/removing tokens from the server.
