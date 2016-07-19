package Controller

import (
    "net/http"
    "finance-go/user-management/Service"
)

func DeleteUser (req *http.Request, w http.ResponseWriter) {
    Service.DeleteUser()
}