package controller

import (
    "github.com/gorilla/mux"
    "finance-go/common/model"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(route.HandlerFunc)
    }
    return router
}

var root = "/api/users"

var routes = model.Routes {
    model.Route {
        "GetUser",
        "GET",
        root + "/{username}",
        GetUser,
    },
    model.Route {
        "AddUser",
        "POST",
        root,
        AddUser,
    },
    model.Route {
        "DeleteUser",
        "DELETE",
        root + "/{username}",
        DeleteUser,
    },
    model.Route {
        "Login",
        "POST",
        root + "/login",
        Login,
    },
    model.Route {
        "Logout",
        "POST",
        root + "/logout",
        Logout,
    },
}