# GoSandBox

My Collection of Code Running in GO ,

Compiler : Gogland by intellij
Synatx: Variable and package are lower and mixed case
        Exported functions and fields have initial upper-case character

Semicolons Not Ended (Usally because compiler add them in for me), that is try

Go would be my default backend API for pretty much every Server

#Highlight in Go
-Async By Default,
-Coroutine , HTTP Server Build -in,
-Good for Middleware,
-MicroService
-Web Development
-Work well with SQL

Go - OOP Thinking
- Define Custom Interface
- Custom types can implement one or more interface
- Custom type can have member methods
- Custom struct(data structure) can have member fields

Go - Doesnt Support
- Type inheritance
- Method or operator overloading
- Structure exception handling
- Implicit numeric conversions , int to flow

Gogland Tricks
Ctrl + Mouse  = Types


-----Working on it

1.	Golang development environment (GoPath, Docker, Appengine)
2.	Golang syntax and language structures(primitives, types, pointers, goroutines )
3.	Golang toolchain (go build, cgo)
4.	Golang package structure and dependencies (go mobile)
5.	Golang web programming, rest api development, Appengine, AWS
6.	Golang system level development
●	static/dynamic linking
●	language bindings (Java/Android, C, Python)
7.	Golang applications
●	command line
●	desktop
●	system

## About Import Package

-- If run into conflict
import {
 la "LoggerX"
 lb "LoggerXY"
}

Calling it would la.LoggerX

-- if want to import as init
import {
 _ "logging"
}

## Go Standard CLI
go build
go clean
go env
go fix
go get
go install
go list
go run
go test
go tool
go version

For Future reference can find in here
https://github.com/hyper0x/go_command_tutorial/blob/master/SUMMARY.md

List of Web Tools and Framework I would use
1. beego(Build HTTP with) https://github.com/astaxie/beego
2. Gogs (Git for yourself)
3. Docker
4. Skynet(Do not forget Terminator) https://github.com/skynetservices
5. NSQ (1 Billions Message per day)https://github.com/nsqio/nsq
6. etcd (Raft)https://github.com/coreos/etcd
7. Groupcache Intent to replace memached in many cases
8. Gobot Do lot of things that I want to do https://gobot.io/ 