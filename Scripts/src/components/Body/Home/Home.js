import React from 'react'
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';

const useStyles = makeStyles(theme => ({
    heroContent: {
        backgroundColor: theme.palette.background.paper,
        padding: theme.spacing(8, 0, 6),
    },
    subContent: {
        padding: theme.spacing(8, 0, 6),
    },
    heroButtons: {
        marginTop: theme.spacing(4),
    },
    subtitleColor: {
        color: theme.palette.secondary.light,
    },
    card: {
        maxWidth: 345,
    },
}));
const techStacks = [
    {
        title: "Golang",
        description: "Simple and fast language with built-in concurrency which is great for backend development",
        image: "/api/public/img/gopher.svg",
        link: "https://golang.org/",
    },
    {
        title: "React",
        description: "Allows easy reuse of components and allows the web application to be dynamic with Javascript",
        image: "/api/public/img/react.svg",
        link: "https://reactjs.org/",
    },
    {
        title: "Material UI",
        description: "Pre-built React components, built with Google's Material Design in mind",
        image: "/api/public/img/material-ui.svg",
        link: "https://material-ui.com/",
    },
    {
        title: "MongoDB",
        description: "A No-SQL database that currently stores all item recipe data and user prices into documents",
        image: "/api/public/img/mongodb.svg",
        link: "https://www.mongodb.com/",
    },
    {
        title: "Docker",
        description: "A containerization application that allows the whole web application to be split into different microservices",
        image: "/api/public/img/docker.svg",
        link: "https://www.docker.com/",
    },
    {
        title: "XIVAPI",
        description: "A RESTful API that has a huge amount of data collected from Final Fantasy XIV",
        image: "/api/public/img/xivapi.png",
        link: "https://xivapi.com/",
    },
]
function Home() {
    const classes = useStyles();
    return (
        <React.Fragment>
            <CssBaseline />
            <div className={classes.heroContent}>
                <Container maxWidth="sm">
                    <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                        FFXIV Profit API
                    </Typography>
                    <Typography variant="subtitle2" align="center" color="secondary" className={classes.subtitleColor}>
                        Made for FFXIV
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
                </Container >
            </div>
            <div className={classes.subContent}>
                <Container>
                    <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                        Built With
                    </Typography>
                    <Typography variant="subtitle2" align="center" color="secondary" className={classes.subtitleColor}>
                        Tech Stack
                    </Typography>
                    <Grid
                        container
                        direction={"row"}
                        justify="space-evenly"
                        alignItems="center"
                    >
                        {techStacks.map(techStack => (
                            <Grid item>
                                <Card className={classes.card}>
                                    <CardActionArea>
                                        <CardMedia>
                                            <img src={techStack.image} height="120" width="345" />
                                        </CardMedia>
                                        <CardContent>
                                            <Typography gutterBottom variant="h5" component="h2">
                                                {techStack.title}
                                            </Typography>
                                            <Typography variant="body2" color="textSecondary" component="p">
                                                {techStack.description}
                                            </Typography>
                                        </CardContent>
                                    </CardActionArea>
                                    <CardActions>
                                        <Button href={techStack.link} target="_blank" rel="noopener" size="small" color="secondary">
                                            Learn More
                                        </Button>
                                    </CardActions>
                                </Card>
                            </Grid>
                        ))}
                    </Grid>
                </Container>
            </div>
        </React.Fragment>
    )

}
export default Home