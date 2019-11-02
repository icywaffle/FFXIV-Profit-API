import React from "react"
import IconButton from "@material-ui/core/IconButton"
import MenuItem from "@material-ui/core/MenuItem"
import { makeStyles } from "@material-ui/core/styles"
import Avatar from "@material-ui/core/Avatar"

import List from "@material-ui/core/List"
import ListItem from "@material-ui/core/ListItem"
import ListItemText from "@material-ui/core/ListItemText"

import ClickAwayListener from "@material-ui/core/ClickAwayListener"
import Grow from "@material-ui/core/Grow"
import Paper from "@material-ui/core/Paper"
import Popper from "@material-ui/core/Popper"
import MenuList from "@material-ui/core/MenuList"

const useStyles = makeStyles((theme) => ({
    root: {
        display: "flex",
    },
    paper: {
        marginRight: theme.spacing(2),
    },
    logout: {
        marginRight: theme.spacing(2),
        [theme.breakpoints.up("sm")]: {
            display: "none",
        },
    },
    avatar: {
        margin: 10,
        width: 24,
        height: 24,
    },
    menu: {
        zIndex: theme.zIndex.drawer + 2,
    },
}))

function LogoutComponent(props) {
    const classes = useStyles
    const [open, setOpen] = React.useState(false)
    const anchorRef = React.useRef(null)

    const handleToggle = () => {
        setOpen(prevOpen => !prevOpen)
    }

    const handleClose = (event) => {
        if (anchorRef.current && anchorRef.current.contains(event.target)) {
            return
        }

        setOpen(false)
    }

    function handleListKeyDown(event) {
        if (event.key === "Tab") {
            event.preventDefault()
            setOpen(false)
        }
    }

    // return focus to the button when we transitioned from !open -> open
    const prevOpen = React.useRef(open)
    React.useEffect(() => {
        if (prevOpen.current === true && open === false) {
            anchorRef.current.focus()
        }

        prevOpen.current = open
    }, [open])

    return (
        <React.Fragment>
            <div className={classes.root}>
                <IconButton
                    ref={anchorRef}
                    aria-controls="menu-list-grow"
                    aria-haspopup="true"
                    onClick={handleToggle}
                >
                    <Avatar alt={props.userinfo.id} src={"https://cdn.discordapp.com/avatars/" + props.userinfo.id + "/" + props.userinfo.avatar + ".gif"} className={classes.avatar} />
                </IconButton>
                <Popper open={open} anchorEl={anchorRef.current} transition disablePortal>
                    {({ TransitionProps, placement }) => (
                        <Grow
                            {...TransitionProps}
                            style={{ transformOrigin: placement === "bottom" ? "center top" : "center bottom" }}
                        >
                            <Paper id="menu-list-grow">
                                <ClickAwayListener onClickAway={handleClose}>
                                    <MenuList autoFocusItem={open} onKeyDown={handleListKeyDown}>
                                        <List className={classes.root}>
                                            <ListItem>
                                                <ListItemText primary="Logged in as" secondary={props.userinfo.username + " #" + props.userinfo.discriminator} />
                                            </ListItem>
                                        </List>
                                        <MenuItem className={classes.logout} onClick={props.Logout}>Logout</MenuItem>
                                    </MenuList>
                                </ClickAwayListener>
                            </Paper>
                        </Grow>
                    )}
                </Popper>
            </div>
        </React.Fragment>
    )
}

export default LogoutComponent