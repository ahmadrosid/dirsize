## Dirsize

Get directory or file size easily.

## Usage
Default folder :
```bash
dirsize
> .: 2.35 MB
```

Custom folder :
```bash
dirsize --dir test
> .: 2.35 MB
```

Custom folder :
```bash
dirsize --dir test1 --dir test2
> ./test1: 2 KB
```

Get file size :
```bash
dirsize --dir ./main
> ./test1: 2.35 MB
> ./test2: 2.35 MB
```

## Install
```bash
go install github.com/ahmadrosid/dirsize
```

## LICENSE
MIT