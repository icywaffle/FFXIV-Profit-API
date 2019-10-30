import React from 'react'
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({
    heroContent: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(8, 0, 6),
    },
    heroButtons: {
        marginTop: theme.spacing(4),
    },
}));

function Home() {
    const classes = useStyles();
    return (
        <div className={classes.heroContent}>
            <Container maxWidth="sm">
                <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                    FFXIV Profit API
            </Typography>
                <Typography variant="h5" align="center" color="textSecondary" paragraph>
                    A RESTful API that stores item recipe information, and also
                    user defined prices and profits.
            </Typography>
                <div className={classes.heroButtons}>
                    <Grid
                        container
                        direction="row"
                        justify="center"
                        alignItems="center">
                        <Grid item>
                            <Button href="/documentation/" variant="contained" color="primary">
                                Documentation
                            </Button>
                        </Grid>
                    </Grid>
                </div>
            </Container>
        </div>
    )

}
export default Home