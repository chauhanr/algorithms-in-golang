# Understanding servers 

In IT the following terms are all synonymous with the web server. 
mux, multiplexer, request router, router, server, servermux, http router, http server etc. all point
to the concept of the web server.

Almost all the web technologies work on the client server architecture and request and responses
both have a same message structure. The web standards are defined by the IETF (internet
engineeering task force) they look into how the net technology must be used and standards. 

Following are request/responses to and from the servers 

**Client Request** 
```
GET /hello.txt HTTP/1.1 

User-Agent: curl/7.16.3 libcurl/7.16.3 OpenSSL/0.9.7 zlib/1.2.3 

HOST: www.example.com 
Accept-Language: en, mi 
```

**Server Respone**

```
HTTP/1.1 200 OK 

Date: Mon, 27 Nov 2019  12:50:00 GMT

Server: Apache 
Last-Modified: Wed, 30 Octoner 2019

ETags: "37beegdner-1568rrfref" 
Accepts-Ranges: bytes 
Content-Length:51 

Vary: Accept-Encoding 
Content-Type: text/plain 

<body> 
```


