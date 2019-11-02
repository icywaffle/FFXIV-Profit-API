import React from "react"
import { CssBaseline } from "@material-ui/core"
import Typography from "@material-ui/core/Typography"


export default function Page(props) {


    return (
        <React.Fragment>
            {props.sections.map((section) => (
                <React.Fragment>
                    <Typography variant="h5" component="h3">
                        {section.title}
                    </Typography>
                    <Typography component="p">
                        {section.paragraph}
                    </Typography>
                </React.Fragment>
            ))}
        </React.Fragment >
    )
}