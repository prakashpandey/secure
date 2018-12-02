# Secure 

Secure is password hashing library in golang on the top of golang's official crypto package

## How to use

- godoc present at <https://godoc.org/github.com/prakashpandey/secure>

- simply follow the below example

### Example

- download using `go get`

    ```bash
        go get github.com/prakashpandey/secure
    ```

- import simply in your project and use it

- example:

    ```bash
    package main

    import (
        "fmt"

        "github.com/prakashpandey/secure"
    )

    func main() {
        password := "Beautiful Day"

        // Generate password hash using SHA512
        hashedPassword, err := secure.Hash(password, secure.SHA512)
        if err != nil {
            fmt.Printf("Error while hashing password: %s", err.Error())
        }

        // Varify password
        isPasswordValid := secure.MatchHash(password, hashedPassword, secure.SHA512)
        fmt.Printf("Password validated ? %t\n", isPasswordValid)
    }

    ```