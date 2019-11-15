import React, { useState, useEffect } from "react"
import OAuth2Payload from "./authkeys/discord.js"
import Login from "./Login"
function OAuth2(props) {
    // Sample Payload
    /*
        client_id: "",
        client_secret: "",
        grant_type: "authorization_code",
        code: "",
        redirect_uri: "",
        scope: "identify",
    */
    const [currentCode] = useState(props.code.location.search.slice(6))
    const [login, setLogin] = useState(localStorage.getItem("user"))
    // For urlencoded forms, we need to manually change our payload
    function encodePayload() {
        var formBody = []
        for (var property in OAuth2Payload) {
            if ({}.hasOwnProperty.call(OAuth2Payload, property)) {
                var encodedKey = encodeURIComponent(property)
                var encodedValue = encodeURIComponent(OAuth2Payload[property.toString()])
                formBody.push(encodedKey + "=" + encodedValue)
            }
        }
        formBody = formBody.join("&")

        return formBody
    }

    // Requests for an access token, so we can access user's info
    function requestAccessToken() {
        var url = "https://discordapp.com/api/oauth2/token"
        OAuth2Payload.code = currentCode
        var encodedPayload = encodePayload()
        fetch(url, {
            method: "POST",

            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            body: encodedPayload,
        })
            .then((response) => response.json())

            // Once we get the data access_token , we can access the User Info and store it in session
            .then((data) => {
                localStorage.setItem("AccessToken", data.access_token)
                const payload = {
                    AccessToken: data.access_token,
                }
                // We also need to log into the API, since our token will expire
                var APIurl = "https://" + window.location.hostname + "/api/userinfo/login/"
                if (window.location.hostname === "localhost") {
                    APIurl = "http://localhost:8080/api/userinfo/login/"
                }
                fetch(APIurl, {
                    method: "POST",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(payload),
                })
                // Then we can go ahead and get user information with this access token.
                fetch("https://discordapp.com/api/users/@me", {
                    headers: {
                        authorization: `${data.token_type} ${data.access_token}`,
                    },
                })
                    .then((response) => response.json())
                    .then((userdata) => {
                        localStorage.setItem("user", JSON.stringify(userdata))
                        setLogin(localStorage.getItem("user"))

                        // Once we"re done getting data, move the user off of the query string.
                        var redirectURL = "https://" + window.location.hostname + "/api/"
                        if (window.location.hostname === "localhost") {
                            redirectURL = "http://localhost:8080/api/"
                        }
                        window.location.href = redirectURL
                    })
            })

    }


    // Only want to request once, and we only want to do it if we made a request with a code query
    useEffect(() => {
        if (currentCode !== "") {
            requestAccessToken()
        }
    }, [])

    return (
        <Login userinfo={JSON.parse(login)} />
    )

}
export default OAuth2