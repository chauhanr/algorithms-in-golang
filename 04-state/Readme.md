# Handling data passing through URL 

Data can be passed to the server using parameters in the URL e.g: 

`http://host.com/values?key1=value1&key2=value2`  

The key1 and key2 can be accessed in the handler function as: 

```
func foo(w http.ResponseWriter, r *http.Request){
    v := r.Formvalue("key1") 
    io.WriteString(w, "Do my search: "+v) 
}

```

## Redirects 

Redirection is about the server moving the request from the original resource to someplace else: 

* 301 - stands for moved permanently. This means that if the client will remember that and send to
  the redirected URL. 
* 303 - see other. Will change to the method Get
* 307 - temporary redirect. Remember the method change. 




