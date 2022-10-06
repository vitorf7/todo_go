# todo_go
Simple Todolist Testing App

Please make sure that you have your `GOPATH` and you `GOROOT`.

`GOPATH` can be anywhere you want your dependencies, etc to go. 
You no longer need to add your projects there since the introduction of `Go modules`.
`GOROOT` on the other hand is a folder called `libexec` which will contain the binary for go.
For example if using something like `homebrew` on a mac and you installed `go` with it and all binaries, etc are installed in `local/opt/` then
your `GOROOT` will be `/usr/local/opt/go/libexec`.

Add these to your shell file (such as `.zshrc`, etc)

```shell
export GOPATH=$HOME/Code/Go
export GOROOT=/usr/local/opt/go/libexec
export GO111MODULE=on

export PATH="$PATH:$GOPATH/bin"
export PATH="$PATH:$GOROOT/bin"
```
