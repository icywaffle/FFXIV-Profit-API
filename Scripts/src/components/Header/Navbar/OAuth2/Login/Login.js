import React from "react"
import LoginComponent from "./LoginComponent"
import LogoutComponent from "./LogoutComponent"
function Login(props) {
    function Logout() {
        localStorage.removeItem("user")
        localStorage.removeItem("AccessToken")
        fetch("https://" + window.location.hostname + "/api/userinfo/logout")
        window.location.href = "https://" + window.location.hostname + "/api/"
    }

    if (props.userinfo) {
        return <LogoutComponent userinfo={props.userinfo} Logout={Logout} />
    } else {
        return <LoginComponent />
    }
}


export default Login