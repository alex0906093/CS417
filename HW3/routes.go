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
        postHandle,
    },
    Route{
        "Get",
        "GET",
        "/Student/getstudent",
        getHandle,
    },
    Route{
        "Delete",
        "DELETE",
        "/Student",
        deleteHandle,
    },
    Route{
        "Update",
        "POST",
        "/Student",
        updateHandle,
    },
    Route{
        "List",
        "GET",
        "/Student/listall",
        listHandle,
    },
}
