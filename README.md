# Yellowpages

A pluggable HTTP services and middleware library.

## Why?

This package is just a wrapper for an http router.

It doesn't do anything novel. It only provides some guardrails and structure for defining http endpoints.

The idea is that you can encapsulate an entire "service" in a single type. Whatever that means to you.

Think a user service.

You might need to:

- Get data for a single user
- Get data for multiple users
- Update user data
- Delete a user

## Getting started

### Installation and initialization

Add the package to your go.mod

```
go get github.com/terminalPoltergeist/yellowpages
```

Import the types and things into your services/routes package.

I recommend dot-importing. This adds the types, variables, and functions in the yellowpages package to the same namespace as the current package.

```go
package routes

import (
    . "github.com/terminalPoltergeist/yellowpages"
)
```

If this causes declaration conflicts in your package, you can just import normally and use the `yellowpages.` prefix.

### Write your first service

```go
package routes

import (
    . "github.com/terminalPoltergeist/yellowpages"
    "github.com/julienschmidt/httprouter"
)

type User struct {
    Name string
    ID string
}

func GetUser(userID string) User {
    // get the user from db
    // create User object
    // return User
}

// note: Service, Endpoint, and Action are types defined in the yellowpages package
var User Service = Service {
    Endpoint {
        Action: GET,
        Path: "/user/:slug",
    } : func (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        userID := p.ByName("slug")

        // get the user by userID
        user := GetUser(userID)

        // marshal the User object to json
        // will look like
        /*
        {
            "Name" : "user",
            "ID" : "id123"
        }
        */
        buffer, err := json.Marshal(&user)
        if err != nil {
            // handle error
        }

        w.Write(buffer)
    }
}
```

What is this doing?

A `Service` is of type `map[Endpoint]httprouter.Handle`.

An `Endpoint` is simply a combination of an HTTP verb, called an `Action`, and a route path.

The `httprouter.Handle` function comes from [github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

### Register the service with your application router

```go
package main

import (
    "github.com/julienschmidt/httprouter"

    "./routes" // import your routes package
)

func main() {
    var router *httprouter.Router = httprouter.New()

    routes.User.Register(router)

    // start the router
}
```
