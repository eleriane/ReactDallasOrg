rm -rf /usr/local/go
wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar -zxvf go1.8.linux-amd64.tar.gz -C /usr/local/

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin