![01-edu-system-blue](https://raw.githubusercontent.com/GoArtyom/study/1a66b22c5b511ccce94b582481a45dfd7f001d3a/alem.svg)
# net-cat

## Description

In this project, I recreated NetCat in a server-client architecture, which can be used in server mode on a given port, listening for incoming connections, or it can be used in client mode, trying to connect to a given port and passing information to the server.

- NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

This program work similar to how the original NetCat works, i.e. it creates a group chat. The project is made according to the following requirements:
- TCP connection between server and multiple clients (relation of 1 to many).
- A name requirement to the client.
- Control connections quantity.
- Clients must be able to send messages to the chat.
- Do not broadcast EMPTY messages from a client.
- Messages sent, must be identified by the time that was sent and the user name of who sent the message, example : `[2020-01-20 15:48:41][client.name]:[client.message]`
- If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
- If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
- If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
- All Clients must receive the messages sent by other Clients.
- If a Client leaves the chat, the rest of the Clients must not disconnect.
- If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: `[USAGE]: ./TCPChat $port`
## Usage

### How to install

```
git clone git@git.01.alem.school:aecheist/net-cat.git && cd net-cat/

go build -o TCPChat cmd/main.go 
```

### How to go run

```
./TCPChat $port 
```
### Example

```console
$ ./TCPCha 2525
Listening on the port :2525
```

Client1 (Yenlik):

```console
$ nc localhost 2525
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Yenlik
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
[2020-01-20 16:04:10][Yenlik]:
Lee has joined our chat...
[2020-01-20 16:04:15][Yenlik]:
[2020-01-20 16:04:32][Lee]:Hi everyone!
[2020-01-20 16:04:32][Yenlik]:
[2020-01-20 16:04:35][Lee]:How are you?
[2020-01-20 16:04:35][Yenlik]:great, and you?
[2020-01-20 16:04:41][Yenlik]:
[2020-01-20 16:04:44][Lee]:good!
[2020-01-20 16:04:44][Yenlik]:
[2020-01-20 16:04:50][Lee]:alright, see ya!
[2020-01-20 16:04:50][Yenlik]:bye-bye!
[2020-01-20 16:04:57][Yenlik]:
Lee has left our chat...
[2020-01-20 16:04:59][Yenlik]:
```

Client2 (Lee):

```console
$ nc localhost 2525
Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: Lee
[2020-01-20 16:03:43][Yenlik]:hello
[2020-01-20 16:03:46][Yenlik]:How are you?
[2020-01-20 16:04:15][Lee]:Hi everyone!
[2020-01-20 16:04:32][Lee]:How are you?
[2020-01-20 16:04:35][Lee]:
[2020-01-20 16:04:41][Yenlik]:great, and you?
[2020-01-20 16:04:41][Lee]:good!
[2020-01-20 16:04:44][Lee]:alright, see ya!
[2020-01-20 16:04:50][Lee]:
[2020-01-20 16:04:57][Yenlik]:bye-bye!
[2020-01-20 16:04:57][Lee]:^C
```

### Audit link

* <a href="https://github.com/alem-01/public/tree/58a361cb2d98e9da144efeecf15435d1fc3bcd6a/subjects/net-cat/audit" target="_blank">net-cat audit</a>

### Autor

* <a href="https://github.com/GoArtyom" target="_blank">GitHub.com/GoArtyom</a>
