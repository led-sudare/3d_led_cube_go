# How to Setup Develop Environment

## 1. Install Golang

Recomended go version: 1.11


### Install Instruction for linux(raspbian)

```shell
wget https://storage.googleapis.com/golang/go1.11.1.linux-armv6l.tar.gz 
sudo tar -C /usr/local -xzf go1.11.1.linux-armv6l.tar.gz
```

add a below to end of ~/.bash_profile  
```shell
export PATH=$PATH:/usr/local/go/bin
```


Referenced from: https://golang.org/doc/install


## 2. Setup Environment Variable
add belows to end of ~/.bash_profile
```shell
....
export "GOPATH=$HOME/go"
export "PATH=$PATH:$GOPATH/bin"
```

then, apply these settings.
```shell
source ~/.bash_profile
```


## 3. Run "go get"s

```shell
go get gonum.org/v1/gonum/mat
go get github.com/jessevdk/go-assets-builder
go get github.com/stretchr/testify
```

## 4. Clone The Code

Get the code
```shell
mkdir -p $GOPATH/src
cd $GOPATH/src
git clone https://github.com/YGFYHD2018/3d_led_cube_go.git
cd 3d_led_cube_go
```

## 5. Build And Run The Program

```shell
go run main.go
```
or if you want executable file.
```shell
go build
./3d_led_cube_go
```
  
This command starts LedFramework whitch can receive "JSON Orders To Show Content" by HTTP(port 8081).  
The target to send "Raw Order To Show Content" by UDP is set "localhost:9001" by default.  
It can be changed by run with "-d" option.  
  

Ex.  
```shell
go run main.go -d 192.168.0.xx:9001
```
or
```shell
go build
./3d_led_cube_go -d 192.168.0.xx:9001
```

## 6. (Optional) Setup Realsense Reciver

install and go get zeromq

```shell
sudo apt-get update
sudo apt-get install -y \
    git build-essential libtool \
    pkg-config autotools-dev autoconf automake cmake \
    uuid-dev libpcre3-dev libsodium-dev valgrind
# only execute this next line if interested in updating the man pages as well (adds to build time):
sudo apt-get install -y asciidoc

git clone git://github.com/zeromq/libzmq.git
cd libzmq
./autogen.sh
# do not specify "--with-libsodium" if you prefer to use internal tweetnacl security implementation (recommended for development)
./configure --with-libsodium
make check
sudo make install
sudo ldconfig
cd ..

git clone git://github.com/zeromq/czmq.git
cd czmq
./autogen.sh && ./configure && make check
sudo make install
sudo ldconfig
cd ..
```

```shell
go get github.com/zeromq/goczmq
```

run or build with '-tags=realsense' option witch enable to receive image data from realsense module.   

Ex.  
```shell
go run main.go -tags=realsense -d 192.168.0.xx:9001
```
or
```shell
go build -tags=realsense
./3d_led_cube_go -d 192.168.0.xx:9001
```


Install instruction of czmq is referenced form
https://github.com/zeromq/czmq#building-and-installing


# Others

## If you add file(s) to under "/asset" ...

You have to run 'build_assets.sh' to generate code that implement go code and binary assets into one binary.  
Please see the detail https://github.com/jessevdk/go-assets-builder 

