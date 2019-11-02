import React from "react"
import ReactDOM from "react-dom"
import Header from "./components/Header"
import Body from "./components/Body"
import { createMuiTheme, ThemeProvider } from "@material-ui/core/styles"

const theme = createMuiTheme({
    palette: {
        type: "dark",
        primary: {
            light: "#484848",
            main: "#212121",
            dark: "#000000",
            contrastText: "#ffffff",
        },
        secondary: {
            light: "#a7c0cd",
            main: "#78909c",
            dark: "#4b636e",
            contrastText: "#000000",
        },
    },
})

ReactDOM.render(
    <ThemeProvider theme={theme}>
        <Header />
        <Body />
    </ThemeProvider>
    ,
    document.getElementById("root")
)


