import React from "react"
import { makeStyles } from "@material-ui/core/styles"
import AppBar from "@material-ui/core/AppBar"
import Button from "@material-ui/core/Button"
import GitHubIcon from "@material-ui/icons/GitHub"
import IconButton from "@material-ui/core/IconButton"
import Grid from "@material-ui/core/Grid"

import MenuIcon from "@material-ui/icons/Menu"

import { BrowserRouter as Router, Route } from "react-router-dom"
import OAuth2 from "./OAuth2"

const useStyles = makeStyles(theme => ({
    button: {
        margin: theme.spacing(1),
    },
    backendButton: {
        margin: theme.spacing(1),
        color: theme.palette.secondary.light,
    },
    appBar: {
        position: "sticky",
    },
}))

export default function Navbar() {
    const classes = useStyles()
    const sections = [
        ["Profits", "/"],
        ["API", "/api/"],
        ["Documentation", "/api/documentation/"],
    ]
    return (
        <AppBar className={classes.appBar} color="default">
            <Grid
                container
                direction="row"
                justify="space-between"
                alignItems="center"
            >
                <Grid item xs={2}>
                    <IconButton className={classes.menuButton} color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                </Grid>
                <Grid item xs={7}>
                    {sections.map((section) => (
                        <Button
                            href={section[1]}
                            className={classes.button}
                        >
                            {section[0]}
                        </Button>
                    ))}
                </Grid>
                <Grid item xs={3}>
                    <Grid container
                        direction="row"
                        justify="flex-end"
                        alignItems="center"
                    >
                        <Grid item>
                            <IconButton href="https://github.com/icywaffle/marketboard-frontend" target="_blank" rel="noopener" className={classes.button} aria-label="Back-End">
                                <GitHubIcon />
                            </IconButton>
                        </Grid>
                        <Grid item>
                            <IconButton href="https://github.com/icywaffle/marketboard-backend" target="_blank" rel="noopener" className={classes.backendButton} aria-label="Back-End" color="secondary">
                                <GitHubIcon />
                            </IconButton>
                        </Grid>
                        <Grid item>
                            <Router>
                                <Route path="/api*" render={(querycode) => <OAuth2 code={querycode} />} />
                            </Router>
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
        </AppBar>
    )
}