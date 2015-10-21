package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Post",
        "POST",
        "/Student",
        postHandle
    }
    Route{
        "Get",
        "GET",
        "/Student/getstudent?{studentID}",
        getHandle
    }
    Route{
        "Delete",
        "DELETE",
        "/Student",
        postHandle
    }
    Route{
        "Update",
        "POST",
        "/Student",
        postHandle
    }
    Route{
        "List",
        "GET",
        "/Student",
        postHandle
    }
}
