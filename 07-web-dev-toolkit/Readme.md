# Web Dev Toolkit 

This is a collection of utilities that we can use in web development. They include topics like: 

1. HMAC (hashing algo to hash sensitive information) 
2. Base64 encoding 
3. Web Storage 
4. Context 
5. TLS & HTTPS
6. JSON handling 


## HMAC - Hash Message Authentication Code 

This util is used to hash message or sensitive information so that we can check whether the use has
indeed changed the sensitive information. The HMAC algorithm needs a private key that is used to
hash the information with and is not shared with the client. Therefore the client does not know the
key and hence cannot change the sensitive information. The concept is very similar to that of a checksum. 

The HMAC utility is present under the crypto package and also uses the sha256 algo. 

``` 
   h := hmac.New(sha256.New, []byte("key")) 
   io.WriteString(h, info)
   enc := fmt.Sprintf("%x", h.Sum(nil))  
```

## Context 
context package in golang was created make it easy to pass request scoped values, cancellation
signals and deadline across API boundaries and to all go routines. 

```

func main() {
   ctx, cancel := context.WithCancel(context.Background()) 
 
   for n := range gen(){
        if n != 5 {
            fmt.Printf("%s\n",n)
        }else{
           cancel() 
           break 
        }
   }
   time.Sleep(1 * time.Minute) 
}

func gen(ctx context.Context) <-chan int{
   ch := make(chan, int) 

   go func(){
        var n int
        n = 0  
        for {
           select{
           case <- ctx.Done(): 
	            return 
           case  ch <- n: 
                  n++ 

           }
        }
   }
}

```

context in the example able prevents the go routine in the gen() function from leaking which will
happen if the `ctx.Done()` is not passed from the infinite for loop. 

The Context Type is defined as: 

```
type Context interface{
    // Done is a channel returned from context and it is closed when the context cancels 
    Done() <- chan struct{} 

    // Err indicates the error about why the context was closed or terminated. 
    Err() error 
  
    // Deadline returns the time when the context will expire. if any. 
    Deadline() (deadline time.Time, ok bool) 
  
    // value returns a value that is set in the context against a key | generic key and value  
    // so you can set literally any golang type as key and value pair.  
    Value(key interface{}) interface{}   
}

```

Context can not only be used in the context of web apps (although it has good amount of applications
there however it canbe used in ohter use cases as well. 


