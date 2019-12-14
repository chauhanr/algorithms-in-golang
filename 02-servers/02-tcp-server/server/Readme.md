# TCP Server connection reading and writing

The code under 02-tcp-server/server does the following: 
1. creates a tcp server and starts it up for receiving connect
2. Once the server is up we can connect to it using telnet client.
3. Once the telnet make a connect we read input from the user and then reprint the same line back. 
4. Finally we add a 10 sec timeout to the connection so that we close the connection to the server. 

Main packages used: 
`bufio.NewScanner(ioReader) *Scanner`- this accepts the read interface and reutrns a scanner object that can
  be used to read form. 
	* Scanner when it is iterated upon will return a line of text at every carriage return and
	  will ends its operations when it hits an error or end of line. 
	* We have to be careful of scanner in the context of a TCP connection as the connection is
	  an open stream and will never give a EOF until terminated.
	* Due to this reason the example above sets a timeout which sends and error and ends the
	  scanner loop. 

`net Connection interface` [net/conn](https://golang.org/pkg/net/#Conn) 
```
  type Conn interface{
        Read(b []byte) (n int, err error) 
        Write(b []byte) (n int, err error) 
        Close() err  
        LocalAddr() Addr
        RemoteAddr() Addr
        SetDeadline(t time.Time) error 
        SetReadDeadline(t time.Time) error 
        SetWriteDeadline(t time.Time) error  
  }
```
Here instead of using the Read and Write method we use the scanner to read and write to the
connection. 

## Listen 

`net.Listen("tcp", "port") Listner, error` - help start a tcp server on port specified also it return a
listener or an error if the sever was not started.

`Listner.Accept() (Conn, error)` - This accepts connections from the server that can be used to
process the connections to the server.   

## Dial 
In order to read from the server we need to use the Dial. the example is under the dial folder. The
reader and dial folder programs need to be used in tandum. 


