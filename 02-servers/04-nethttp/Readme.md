# net/http package 

HTTP is a protocol that is built on top of the tcp protocol. Therefore the net/http package has a
lot of functionality where the server or http multiplexers are already built for us. 


## http.Handler
Any golang type that needs to serve http request must implement the Handler interface 

```
type Handler interface{
    ServeHTTP(http.ResponseWriter, *http.Request) 
}
```

## Server 

```
http.ListenAndServe(add string, handler http.Handler) error


http.ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error 
```

## Request

```
type Request struct{

	Method string
	URL    *url.URL 
 	/*
	    Header = map[string][]string 
  	    "Accept-Encoding": {"gzip, deflate"}, 
            "Accept-Language": {"en-US"}, 
	*/
        Header Header 
        Body   io.ReadCloser 
 	ContentLenght int64 
  	Host   string 
	Form   url.Values 
	PostForm url.Values 
	MultipartForm  *mutlipart.Form 
        /* GetBody is a function that need to be called when the Body needs read multiple times*/
        GetBody func() (io.ReadCloser, error)
        Close bool 
        Host string // this indicates where the host machine where the request originated. 
        RemoteAddr string  // this is a the IP:Port of the machine from which the request origs */ 
        Cancel <-chan struct{} 
        Response *Response  // this is populated during client redirects.   
}
```

## ResponseWriter 

```
type ResponseWriter interface{
   /*
	Returns the header that are part of the response object that is to be sent back to the 
	client. Changing the header map after a call to WriteHeader (or write) has no effect. 
    */
   Header() Header
   /*
	Write the data to the connection as part of the HTTP
	If WriteHeader has not been called Write() will set the http.StatusOK as the status before
	writing the response body. 
	If Content-Type is not set explicitly then the content's 512 bytes are taken to determine
	the content type by passing it to DetectContentType method. If the data written to the 
	body is few KB there is no flush call. 
   */
   Write([]byte) (int, error) 
  /* 
    	Write Header - sets the http.StatusCode to the status code. Explicit calls to this method will 
   	Set the status code if however it is not called the first write will set it to StatusOK 
	HTTP code 1xx-5xx must be set on the header status. 
  */ 
   WriteHeader(status int)  
}

```

Because of the `Write([]byte) (int, error)` method the ReponseWriter is a writer type and therefore we can
use the methods to write to the Response object. 
* `fmt.Fprintf(w io.Writer, string,inteface{} ... )` or `fmt.Fprintln(w,string)`
* `io.WriteString(io.Writer, string)` 




