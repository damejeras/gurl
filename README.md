# gurl

A tool to parse URL written in Go

```bash
$ gurl --help            
Usage:
  gurl [OPTIONS] input

Application Options:
  -f, --format= format output with text/template syntax

Help Options:
  -h, --help    Show this help message
```

### Usage examples
```bash
$ gurl --format "{{(index .Query.scopes 0)}}" "/test?scopes=a%20b"
a b

$ gurl --format "{{.Scheme}}://{{.Host}}{{.Path}}" "https://example.com/test?scopes=a%20b"
https://example.com/test
```
Can be also used in conjunction with `xargs`.