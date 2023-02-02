# ezAES

A tool copy from sprig to encrypt string easily with AES.

## Introduction

When I was trying to build a pipeline with [helm chart](helm.sh/), I noticed that it's a little difficult to generate a AES string with command line tool like `openssl`. So I dig into [sprig](https://github.com/Masterminds/sprig) and found there're [some go-implemented functions](https://github.com/Masterminds/sprig/blob/581758eb7d96ae4d113649668fa96acc74d46e7f/crypto.go#L597).

To easily encrypt and decrypt a string with AES. I write this simple tool.

Of course, you can this tool to encrypt any string or just debug helm chart like me.

## Usage

Download ezAES or build yourself.

```
# interactively
ezAES
# encrypt a string
ezAES -t hello
# decrypt a string
ezAES -d -t sUb5DTeg6ysAxdelQu+R3Za1ZI0cwRzumB25HtHHSRM=
# unsafely encrypt a string
ezAES -t hello -k 123
```

## LICENSE

MIT
