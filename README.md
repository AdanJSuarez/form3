# Form3 Homework

Form3 homework.

Ref:

https://github.com/form3tech-oss/interview-accountapi

https://github.com/form3tech-oss/go-form3

## Client library

The library is a Go library to be use as so. For simplicity, it implements only limited functionality of the Account endpoints of Form3.

For accounts, it implements:
- `Create` a new bank account.
- `Fetch` a bank account.
- `Delete` a bank account.

A nice future feature we could include (not included) is to expose a C-shared API that would allow the use of this library from any other language than Go that can read C libraries, like Java, Python, Rust, C++, C, etc.

## Use

The library use an extension of the principle of [Hiding Information](https://en.wikipedia.org/wiki/Information_hiding) therefore you only can reach the functionality expose in form3.go that would be:
```go
    type Form3 struct
```

```go
    type Account interface {
        Create()
        Fetch()
        Delete()
    }
```

## Vendor

The eternal discussion of pushing the dependency to your repo. I saw in your client [repo](https://github.com/form3tech-oss/go-form3) you included it, so I took the same approach.
This, as everything in live, brings pros and cons.
This blog post talk about this... and he is a big fun btw :) [link](https://blog.boot.dev/golang/should-you-commit-the-vendor-folder-in-go/)


## Network
https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach
