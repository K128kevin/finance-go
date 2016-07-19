package Controller

import (
    "github.com/gorilla/mux"
    "finance-go/Common/Model"
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

var routes = Model.Routes {
    Model.Route {
        "GetUser",
        "GET",
        root + "/{username}",
        GetUser,
    },
    Model.Route {
        "AddUser",
        "POST",
        root,
        AddUser,
    },
    Model.Route {
        "DeleteUser",
        "DELETE",
        root + "/{username}",
        DeleteUser,
    },
    Model.Route {
        "Login",
        "POST",
        root + "/login",
        Login,
    },
    Model.Route {
        "Logout",
        "POST",
        root + "/logout",
        Logout,
    },
}