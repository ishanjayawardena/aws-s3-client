# AWS S3 Reader

An AWS S3 command line tool that reads content at a specified s3 key and displays it in the operating system's default pager.

### Howe to build
```bash
make build
```
`make build` will generate four binaries one for each application env/AWS region combination.

`Makefile` contains other possible targets such as `install` for installing the bins to your local `$GOPATH` and `clean` for cleaning up binaries.

### Usage
```bash
stg-s3-eu -k <S3-Object-key> # without bucket name
```

Optional arguments
```bash
d 10ms # S3 read timeout in timeunit notaion
```