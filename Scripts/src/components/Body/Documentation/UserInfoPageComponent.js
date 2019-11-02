import React from "react"

export default function UserInfoPageComponent() {
    return [
        {
            title: "User Information",
            paragraph:
                <div>
                    <p>
                        User information is based off of whether the backend can use the Access Token, to obtain Discord User information or not.
                        If the access token cannot provide information, it must have been a bad access token.
                    </p>
                    <p>
                        We can basically secure a user's storage, based on the access token.
                        If the access token was successful, we basically can now store the user's session to the RESTful API.
                    </p>
                    <p>
                        Once a user's session has been established, you can now access your own user storage.
                    </p>
                    <p>
                    <strong>Note:</strong> You can only access this endpoint if you're logged in.
                    </p>
                </div>,
        },
        {
            title: "GET /userinfo/recipe/:recipeid",
            paragraph: 
            <div>
                <p>
                GET {window.location.protocol}//{window.location.hostname}/api/userinfo/recipe/33180
                </p>
                <p>
                    With a given user info ID, you can obtain your own recipes that are stored in your user storage.
                </p>
            </div>,
        },
        {
            title: "POST /userinfo/",
            paragraph:
            <div>
                <p>
                POST {window.location.protocol}//{window.location.hostname}/api/userinfo/
                </p>
            </div>,
        },
    ]
}