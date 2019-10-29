import React from "react"
import {
    BrowserRouter as Router,
    Route,
} from "react-router-dom"
import Home from "./Home"

function HomePage() {
    return <Home />
}


function Body() {
    return (
        <Router>
            <Route path="/" exact component={HomePage} />
        </Router>
    )
}

export default Body