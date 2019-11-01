import React from "react"

export default function IntroPageComponent() {
    return [
        {
            title: "Introduction",
            paragraph:
                <div>
                    <p>
                        Welcome to the Documentation page for the RESTful API endpoint for FFXIV Profit!
                    This API is under heavy development, and it's really only meant for FFXIV Profit.
                    Most of the information about recipes you can get, are from XIVAPI, and they're just stored here
                    for convenience.
                    </p>
                    <p>
                        The important content comes from the prices that all users generate when
                    they put their information inside the profits page. This generates all the prices and profits
                    for a specific recipe. With this, it populates your personal profits over time. Eventually, we would be able to
                    get an overall profit for a specific item, or even more later on, we can implement graphs and statistics for an item.
                    </p>
                </div>,
        },
        {
            title: "API Access",
            paragraph:
                <div>
                    <p>
                        There aren't any rate-limiting methods applied currently. There will be, but just not now.
                        It would be limited to IP, with no Access Token, since that would be the simplest to implement.
                    </p>
                </div>,
        },
        {
            title: "Getting Started",
            paragraph:
                <div>
                    <p>
                        To get started, go ahead and check the available open endpoints in the menu.
                        The only methods that are available are GET and POST requests.
                    </p>
                </div>,
        },
    ]

} 