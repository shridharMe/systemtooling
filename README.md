# System tooling with Go
Inspired by the linuxacademy training lesson on System Tooling with Go
https://linuxacademy.com/cp/modules/view/id/291

```sh 
#install cobra
go get -u github.com/spf13/cobra/cobra

#create cobra project for command line tool
cobra init --pkg-name github.com/shridharMe/systemtooling/motd shridharMe/systemtooling/motd

# to run the project

go get github/shridharMe/systemtooling
cd github.com/shridharMe/systemtooling
go run motd/main.go

```
