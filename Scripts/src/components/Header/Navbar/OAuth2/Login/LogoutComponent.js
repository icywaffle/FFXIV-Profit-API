import React from "react"
import IconButton from "@material-ui/core/IconButton"
import Menu from "@material-ui/core/Menu"
import MenuItem from "@material-ui/core/MenuItem"
import { makeStyles } from '@material-ui/core/styles'
import Avatar from "@material-ui/core/Avatar"

import List from "@material-ui/core/List"
import ListItem from "@material-ui/core/ListItem"
import ListItemText from "@material-ui/core/ListItemText"

const useStyles = makeStyles(theme => ({
    root: {
        width: '100%',
        maxWidth: 360,
        backgroundColor: theme.palette.background.paper,
    },
    logout: {
        marginRight: theme.spacing(2),
        [theme.breakpoints.up('sm')]: {
            display: 'none',
        },
    },
    avatar: {
        margin: 10,
        width: 24,
        height: 24,
    },
}))

function LogoutComponent(props) {
    const classes = useStyles
    const [anchorEl, setAnchorEl] = React.useState(null)

    const handleClick = event => {
        setAnchorEl(event.currentTarget)
    }
    const handleClose = () => {
        setAnchorEl(null)
    }
    return (
        <React.Fragment>
            <IconButton aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
                <Avatar alt={props.userinfo.id} src={"https://cdn.discordapp.com/avatars/" + props.userinfo.id + "/" + props.userinfo.avatar + ".gif"} className={classes.avatar} />
            </IconButton>
            <Menu
                id="simple-menu"
                anchorEl={anchorEl}
                anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
                transformOrigin={{ vertical: 'top', horizontal: 'right' }}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
            >
                <List className={classes.root}>
                    <ListItem>
                        <ListItemText primary="Logged in as" secondary={props.userinfo.username + " #" + props.userinfo.discriminator} />
                    </ListItem>
                </List>
                <MenuItem className={classes.logout} onClick={props.Logout}>Logout</MenuItem>
            </Menu>
        </React.Fragment>
    )
}

export default LogoutComponent