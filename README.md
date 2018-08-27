# Ring [WIP]
Consistent hashing paper implementation using Red Black Tree 
![ring](http://paperplanes-assets.s3.amazonaws.com/consistent-hashing.png)

## Example Usage

```go
ring:=NewRing([]string{"server-1","server-2","server-3"},1)
node:=ring.Get("foo")
```


## TODO

- Test cases
- Performance test for xxhash

Ring Image source: http://paperplanes-assets.s3.amazonaws.com/consistent-hashing.png
