import React from "react"

import LeftNavComponent from "./LeftNavComponent"
import RightNavComponent from "./RightNavComponent"

function Routes() {
    return (
        <div
            uk-sticky="animation: uk-animation-slide-top; sel-target: .uk-navbar-container; cls-active: uk-navbar-sticky; cls-inactive: uk-navbar-transparent; top: 200"
        >
            <nav className="uk-navbar-container uk-background-default uk-navbar uk-dark">
                <LeftNavComponent />
                <RightNavComponent />
            </nav>
        </div>

    )
}

export default Routes