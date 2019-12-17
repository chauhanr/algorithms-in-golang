# Routing 

ServeMux is a router that we can use to make the routing possible. It matches the url of each
incoming request to a list of already registered patterns and invokes appropriate handlers for that
pattern. 

Patterns are matched against: 
1. fixed rooted paths like '/favicon.ico' or 
2. rooted subtree '/images/' 

Longer pattern always takes precedence over shorter pattern so if we have a handler againt both
'/images/' and '/images/thumbnail/' then the latter handler will be called for request that start
with 'images/thumbnail/'. 
Note: since the path ending is a slash at the end is a rooted subtree the pattern "/" matches all
paths not matched by other registered patterns, not just the URL '/'. 

Lets look at a few examples: 
* `/dog/`  will catch not only /dog path but also /dog/something or /dog/something/else etc. unless
  you have a different handler mapped to those URL. 
* however `/cat` will only map the /cat and will not map /cat/some etc. 


ServeMux is an empty struct 

```
type ServeMux struct{
    
}
``` 
The following are the implementation methods associated with servemux 

```
func NewServeMux() *http.ServeMux 
func (mux *ServeMux) Handle(pattern string, handler http.Handler) 
func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) 
func (mux *ServeMux) Handler(r *http.Request) (h http.Handler, pattern string) 
func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) 

```



