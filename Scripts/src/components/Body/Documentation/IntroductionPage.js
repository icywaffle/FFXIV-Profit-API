import React from "react"
import { CssBaseline } from "@material-ui/core"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"

const Sections = [
    {
        title: "Introduction",
        paragraph: `Welcome to the Documentation page for the RESTful API endpoint for FFXIV Profit!
            This API is under heavy development, and it's really only meant for FFXIV Profit.
            Most of the information about recipes you can get, are from XIVAPI, and they're just stored here
            for convenience.`,
    },
    {
        title: "API Access",
        paragraph: ` There aren't any rate-limiting methods applied currently. There will be, but just not now.`,
    },
    {
        title: "Getting Started",
        paragraph: `To get started, go ahead and check the available open endpoints in the menu.
            The only methods that are available are GET and POST requests.`,
    },
]

export default function IntroductionPage() {

    return (
        <React.Fragment>
            <CssBaseline />
            <Grid container
                direction="column"
                justify="space-evenly"
                alignItems="stretch"
                spacing={5}
            >
                {Sections.map((section) => (
                    <Grid item>
                        <Typography variant="h5">
                            {section.title}
                        </Typography>
                        <Typography variant="body2">
                            {section.paragraph}
                        </Typography>
                    </ Grid>
                ))}
            </Grid>
        </React.Fragment >
    )
}