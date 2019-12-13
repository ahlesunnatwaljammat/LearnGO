set GOPATH=%cd%

set GOBIN=%cd%/bin

set PATH=%PATH%;%GOROOT%\bin

cd src/examples

dep ensure

cd %GOPATH%