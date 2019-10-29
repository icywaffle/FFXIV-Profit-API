import React from "react"

function LeftNavComponent() {
    return (
        <div className="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
            <ul className="nav">
                <li>
                    <a href="/">Home</a>
                </li>
                <li>
                    <a href="/search/">Search</a>
                </li>
            </ul>
        </div>
    )
}

export default LeftNavComponent