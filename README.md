# inmem-db
In Memory DB using HashTable Algorithm

Supports Memory Compression and Fast Read, Write and Delete Operations to Serve as a Memory Cache or Short Term Memory DB.

Takes in The following CLI flags:

### --compress: 

Takes in A Boolean Value(defaults to false) , Set to Enable Memory Compression of Data

### --size: 
Defaults to 1000, Takes in an Integer to Determine No. of Elements allowed in Memory,if Memory is Limited.


## How to Use:

inmem-db uses a HTTP server powered by Fiber to take in Write, Read and Delete Requests in the Memory Cache/DB


**/del (DELETE METHOD)**: Route for deleting data, requires key (e.g. {key: "abrabac"} ) in JSON request

**/get (POST METHOD)**: Route for getting data, requires key (e.g. {key: "abrabac"} ) in JSON request

**/post (POST METHOD)**: Route for getting data, requires key and value (e.g. {key: "abraba" , value: "2020 Saga happened"} ) in JSON request
