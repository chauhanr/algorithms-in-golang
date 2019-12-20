# File Server 

`http.FileServer` - is a utility package that exposes a file system on the system on which the
server is running and presents it on the web. 

The `FileHandler` func takes a `http.FileSystem` interface which implements access to a collection
of named files. 

```
type FileSystem interface{
   Open(name string) (File, error) 
}
```

To use the filesystem on the server we need to provide the path to where we need access. 

```
  http.Handle("/", http.FileServer(http.Dir("/usr/shared/doc") 
```

The http.Dir is a type string that implements the FileSystem interface. 
