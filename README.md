# {{ $base }} 

```shell
# containerly create test0001 --template=https://github.com/thamilton-containerly/containerly-go-grpc-template/archive/refs/heads/main.zip
containerly create {{ $base }} --template [template]
...
```

```shell
containerly build
...
```

```shell
./script/install
```

```shell
containerly deploy --target local
...
```
