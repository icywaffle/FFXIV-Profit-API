import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import GitHubIcon from '@material-ui/icons/GitHub';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';

const useStyles = makeStyles(theme => ({
    button: {
        margin: theme.spacing(1),
    },
}));

export default function Navbar() {
    const classes = useStyles();
    const sections = [
        ["Profits", "/"],
        ["API", "/api/"],
        ["Documentation", "/documentation/"],
    ]
    return (
        <AppBar position="static" color="default">
            <Grid
                container
                direction="row"
                justify="space-between"
                alignItems="center"
            >
                <div>
                    {sections.map(section => (
                        <Button
                            href={section[1]}
                            className={classes.button}
                        >
                            {section[0]}
                        </Button>
                    ))}
                </div>
                <div>
                    <IconButton href="https://github.com/icywaffle/marketboard-frontend" target="_blank" rel="noopener" className={classes.button} aria-label="Back-End">
                        <GitHubIcon />
                    </IconButton>
                    <IconButton href="https://github.com/icywaffle/marketboard-backend" target="_blank" rel="noopener" className={classes.button} aria-label="Back-End" color="secondary">
                        <GitHubIcon />
                    </IconButton>
                </div>
            </Grid>
        </AppBar>
    );
}