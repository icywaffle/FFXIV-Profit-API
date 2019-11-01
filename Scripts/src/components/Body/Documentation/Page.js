import React from "react"
import { CssBaseline } from "@material-ui/core"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"

export default function Page(props) {

    return (
        <React.Fragment>
            <CssBaseline />
            <Grid container
                direction="column"
                justify="space-evenly"
                alignItems="stretch"
                spacing={5}
            >
                {props.Sections.map((section) => (
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