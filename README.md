# Ring
Consistent hashing paper implementation using Red Black Tree 
![ring](http://paperplanes-assets.s3.amazonaws.com/consistent-hashing.png)

## Example Usage

```go
ring:=NewRing([]string{"server-1","server-2","server-3"},1)
node:=ring.Get("foo")
```


## TODO

- More test cases
- Performance test for xxhash

## Paper
https://www.akamai.com/es/es/multimedia/documents/technical-publication/consistent-hashing-and-random-trees-distributed-caching-protocols-for-relieving-hot-spots-on-the-world-wide-web-technical-publication.pdf


Ring Image source: http://paperplanes-assets.s3.amazonaws.com/consistent-hashing.png
