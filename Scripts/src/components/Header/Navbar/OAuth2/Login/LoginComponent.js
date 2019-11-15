import React from "react"
import Button from "@material-ui/core/Button"
function LoginComponent() {
    var discordURL = "https://discordapp.com/api/oauth2/authorize?client_id=598247290972667904&redirect_uri=https%3A%2F%2Fffxivprofit.com%2Fapi%2Fdiscordcode%2F&response_type=code&scope=identify"
    if (window.location.hostname === "localhost") {
        discordURL = "https://discordapp.com/api/oauth2/authorize?client_id=598247290972667904&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fapi%2Fdiscordcode%2F&response_type=code&scope=identify"
    }

    return (
        <Button href={discordURL} >
            Login
        </Button>

    )
}

export default LoginComponent