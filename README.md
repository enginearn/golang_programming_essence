# Go

## Create project

``` PowerShell
mkdir my-app
cd my-app
go mod init my-app
touch main.go

go run main.go
```

## gopherdata/gophernotes

``` PowerShell
# go env GOPATH > $env:GOPATH
mkdir $env:GOPATH\src\github.com\gopherdata
cd $env:GOPATH\src\github.com\gopherdata
git clone https://github.com/gopherdata/gophernotes
cd gophernotes
go install
```

``` PowerShell
mkdir $env:APPDATA\jupyter\kernels\gophernotes
xcopy $env:GOPATH\src\github.com\gopherdata\gophernotes\kernel $env:APPDATA\jupyter\kernels\gophernotes /s
```
