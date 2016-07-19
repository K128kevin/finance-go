package Controller

import (
    "net/http"
    "FinanceGo/UserManagement/Service"
)

func GetUser (req *http.Request, w http.ResponseWriter) {
    Service.GetUser()
}

func AddUser (req *http.Request, w http.ResponseWriter) {
    Service.AddUser()
}

func DeleteUser (req *http.Request, w http.ResponseWriter) {
    Service.DeleteUser()
}

func Login (req *http.Request, w http.ResponseWriter) {
    Service.Login()
}

func Logout (req *http.Request, w http.ResponseWriter) {
    Service.Logout()
}