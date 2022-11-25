# godo


Build instructions for Windows  
https://stackoverflow.com/questions/41566495/golang-how-to-cross-compile-on-linux-for-windows
`GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build`   

