# yotpo-socials-simulator

To run server please config go environment:
1) Add this to .bash_profile:
export GOPATH=$HOME/Development/go-workspace
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
2)Run those commands from terminal
source ~/.bash_profile
brew install go
cd ~/Development
mkdir go-workspace
mkdir   go-workspace/bin
mkdir   go-workspace/pkg
mkdir   go-workspace/src
cd  go-workspace/src
git clone https://github.com/YotpoLtd/yotpo-socials-emulator.git
cd yotpo-socials-emulator
3) Run facebook server with this command: go run yotpo-facebook/Main.go 
4) Open browser and link to http://localhost:8080/me/permissions
5) You'll see this responsce
{
data: [
{
permission: "user_friends",
status: "granted"
},
{
permission: "email",
status: "granted"
},
.........
}
