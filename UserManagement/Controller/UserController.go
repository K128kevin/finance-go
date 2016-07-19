package Controller

import (
    "net/http"
    "finance-go/UserManagement/Service"
)

func DeleteUser (req *http.Request, w http.ResponseWriter) {
    Service.DeleteUser()
}