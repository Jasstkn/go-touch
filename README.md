# go-touch
Tiny program to create a file with all non-existing parent directories

## Installation

You can download executable binaries in [Releases](https://github.com/Jasstkn/go-touch/releases) or use one of the manual methods below:

```bash
# installs to ~/go/bin
go install github.com/Jasstkn/go-touch

# build by yourself
git clone https://github.com/Jasstkn/go-touch.git
go build
./go-touch
```

## Usage

```bash
# creates directories: tmp/, new/ and file.txt
go-touch /tmp/new/file.txt
```
