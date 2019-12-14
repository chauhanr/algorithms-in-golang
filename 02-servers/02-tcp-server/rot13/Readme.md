# Rot 13 server 

The Rot 13 server will take a string and then rotate 13 positions. This is a rudimentary encryption
server 

`go run main.go` This will start the server that will look for a connection. 

We can then call this program using the dialer go to the ../dial and run the main.go with -rw option
as write and message
`go run main.go -rw w -msg <msg>`
