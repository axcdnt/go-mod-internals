This repo is _merely_ a function, exported as it was a library. Its main purpose is to explore Go modules internals.
Modules is the most [recent](https://blog.golang.org/using-go-modules) way of managing dependencies in Go. From version 1.13+ modules became the default.

This repo was first made private. Now it's public so you can use it to learn more. I hope it works for your private repos too. I believe there are different ways to solve this problem. :)

## Resolving private dependencies

### Configurations

#### Environment variable

`export GOPRIVATE="github.com/<username>,github.com/<username>,github.com/<username>"`

Notice that it allows you to define multiple repos.

#### Enforcing SSH via Git

`git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"`

Now, if you want to use it in another project, run the following:

`go get -v github.com/axcdnt/go-mod-internals`

The go.mod must be correctly updated with the code from this lib.

Dependencies resolved, it's now possible to import and use it:

```
package main

import "github.com/axcdnt/go-mod-internals/stringer"

func main() {
	stringer.Tokenize("what a cool string") // ["what", "a", "cool", "string"]
}

```

For more details, please check [Go 1.13](https://golang.org/doc/go1.13#modules).
