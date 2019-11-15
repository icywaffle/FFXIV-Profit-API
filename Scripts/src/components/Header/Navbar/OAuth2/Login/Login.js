import React from "react"
import LoginComponent from "./LoginComponent"
import LogoutComponent from "./LogoutComponent"
function Login(props) {
    function Logout() {
        localStorage.removeItem("user")
        localStorage.removeItem("AccessToken")
        var APIurl = "https://" + window.location.hostname + "/api/userinfo/logout"
        var redirectURL = "https://" + window.location.hostname + "/api/"
        if (window.location.hostname === "localhost") {
            APIurl = "http://localhost:8080/api/userinfo/logout"
            redirectURL = "http://localhost:8080/api/"
        }
        fetch(APIurl, {
            credentials: "include",
        })
        window.location.href = redirectURL
    }

    if (props.userinfo) {
        return <LogoutComponent userinfo={props.userinfo} Logout={Logout} />
    } else {
        return <LoginComponent />
    }
}


export default Login