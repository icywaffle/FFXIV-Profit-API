import React from "react"
import Button from "@material-ui/core/Button"
function LoginComponent() {

    return (
        <Button href={"https://discordapp.com/api/oauth2/authorize?client_id=598247290972667904&redirect_uri=https%3A%2F%2F"
            + window.location.hostname
            + "%2Fapi%2Fdiscordcode%2F&response_type=code&scope=identify"} >
            Login
        </Button>

    )
}

export default LoginComponent