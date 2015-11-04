# yotpo-socials-simulator


To run server please config go environment:<br>
1) Add this to .bash_profile:<br>
export GOPATH=$HOME/Development/go-workspace<br>
export GOROOT=/usr/local/opt/go/libexec<br>
export PATH=$PATH:$GOROOT/bin<br>
export PATH=$PATH:$GOPATH/bin<br>
2)Run those commands from terminal<br>
source ~/.bash_profile<br>
brew install go<br>
cd ~/Development<br>
mkdir go-workspace<br>
mkdir   go-workspace/bin<br>
mkdir   go-workspace/pkg<br>
mkdir   go-workspace/src<br>
cd  go-workspace/src<br>
git clone https://github.com/YotpoLtd/yotpo-socials-emulator.git<br>
cd yotpo-socials-emulator<br>
3) Run facebook server with this command: go run yotpo-facebook/Main.go <br>
4) Open browser and link to http://localhost:8080/me/permissions<br>
5) You'll see this responsce<br>
{<br>
data: [<br>
{<br>
permission: "user_friends",<br>
status: "granted"<br>
},<br>
{<br>
permission: "email",<br>
status: "granted"<br>
},<br>
.........<br>
}<br>
