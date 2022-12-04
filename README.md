this sample code generate dynamic website like React.js

###Running Code
GopherJS [requires Go 1.18 or newer](https://github.com/gopherjs/gopherjs/blob/master/doc/compatibility.md#go-version-compatibility). If you need an older Go
version, you can use an [older GopherJS release](https://github.com/gopherjs/gopherjs/releases).

Install GopherJS with `go install`:

```
go install github.com/gopherjs/gopherjs@v1.18.0-beta1  # Or replace 'v1.18.0-beta1' with another version.
```

Transpile golang to js with the following

```
gopherjs build [module|file]
```

Server with gopherjs 

```
gopherjs serve
```
this command run local server on port 8080

for this sample code you must run the following command

```
gopherjs build main.go
gopherjs server
```

then open your browser and type following address
```
localhost:8080/index.html
```