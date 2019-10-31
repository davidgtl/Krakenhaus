import {MuiThemeProvider} from "@material-ui/core";
import {loginTheme} from "../scenes/common";
import Button from "@material-ui/core/Button";
import Paper from "@material-ui/core/Paper";
import React from "react";
import {logout, roleProp} from "../login/session";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import MenuIcon from '@material-ui/icons/Menu';
import {makeStyles} from '@material-ui/core/styles';
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import Cookies from 'js-cookie'

const useStyles = makeStyles(theme => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
}));

export default function Navbar({props}) {
    const classes = useStyles();
    const [anchorEl, setAnchorEl] = React.useState(null);

    const handleClick = event => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        return () => {
            setAnchorEl(null);
        };
    };

    const handleViewChange = (view) => {
        return () => {
            Cookies.set("view", view);
            window.location.reload();
        };
    };

    const navitems = roleProp("navitems");
    console.log(navitems);

    const menuItems = Object.getOwnPropertyNames(navitems).map(x => <MenuItem onClick={handleViewChange(navitems[x])}>{x}</MenuItem>);

    return (
        <MuiThemeProvider theme={loginTheme}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu"
                                aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
                        <MenuIcon/>
                    </IconButton>
                    <Menu
                        id="simple-menu"
                        anchorEl={anchorEl}
                        keepMounted
                        open={Boolean(anchorEl)}
                        onClose={handleClose}
                    >
                        {menuItems}
                    </Menu>
                    <Typography variant="h6" className={classes.title}>
                        {roleProp("greeting")}
                    </Typography>
                    <Button variant="contained" color="secondary" onClick={() => logout()}>
                        Logout
                    </Button>
                </Toolbar>
            </AppBar>
        </MuiThemeProvider>)
};