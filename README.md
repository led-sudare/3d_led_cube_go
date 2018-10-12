# How to Setup Develop Environment

## 1. Install Golang

Recomended go version: 1.11

## 2. Setup Environment Variable
add belows to end of ~/.bash_profile
```shell
....
export "GOPATH=$HOME/go"
export "PATH=$PATH:$GOPATH/bin"
```

## 3. Run "go get"s

```shell
$ go get gonum.org/v1/gonum/mat
$ go get github.com/jessevdk/go-assets-builder
$ go get github.com/stretchr/testify
```

## 4. Clone The Code

Get the code
```shell
$ cd $GOPATH/src
$ git clone https://github.com/YGFYHD2018/3d_led_cube_go.git
$ cd 3d_led_cube_go
```

## 5. Build And Run The Program

```shell
$ go run main.go
```
  
This command starts LedFramework witch can receive "JSON Orders To Show Content" by HTTP(port 8081).  
The target to send "Raw Order To Show Content" by UDP is set "localhost:9001" by default.  
It can be changed by run "go run main.go" with "-d" option.  
  

Ex.  
```shell
$ go run main.go -d 192.168.0.xx:9001
```


# Others

## If you add file(s) to under "/asset" ...

You have to run 'build_assets.sh' to generate code that implement go code and binary assets into one binary.  
Please see the detail https://github.com/jessevdk/go-assets-builder 

