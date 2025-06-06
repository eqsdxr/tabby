[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

# Tabby

A tiny library for super simple Golang tables

**Get Tabby**

```go
go get github.com/cheynewallace/tabby
```

**Import Tabby**

```go
import "github.com/cheynewallace/tabby"
```

Tabby is a tiny (around 70 lines of code) efficient library for writing extremely simple table based terminal output in Golang.

Many table libraries out there are overly complicated and packed with features you don't need. If you simply want to write clean output to your terminal in table format with minimal effort, Tabby is for you.

Typical examples

- Writing simple tables with heading and tab spaced columns
- Writing log lines to the terminal with evenly spaced columns

## Example With Heading

```go
t := tabby.New()
t.AddHeader("NAME", "TITLE", "DEPARTMENT")
t.AddLine("John Smith", "Developer", "Engineering")
t.Print()
```

**Output**

```
NAME        TITLE      DEPARTMENT
----        -----      ----------
John Smith  Developer  Engineering
```

## Example Without Heading

```go
t := tabby.New()
t.AddLine("Info:", "WEB", "Success 200")
t.AddLine("Info:", "API", "Success 201")
t.AddLine("Error:", "DATABASE", "Connection Established")
t.Print()
```

**Output**

```
Info:   WEB       Success 200
Info:   API       Success 201
Error:  DATABASE  Connection Established
```

## Example With Custom tabWriter

```go
w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
t := tabby.NewCustom(w)
```

## Full Example

- Default writer (os.Stdout) example:

```go
package main

import "github.com/cheynewallace/tabby"

func main() {
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.Print()
}
```

- File writer example:

```go
package main

import (
	"os"
	"text/tabwriter"

	"github.com/cheynewallace/tabby"
)

func main() {
	fd, _ := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer fd.Close()

	w := tabwriter.NewWriter(fd, 0, 0, 4, ' ', 0)
	t := tabby.NewCustom(w)
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.Print()
}
```
