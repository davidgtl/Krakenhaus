import {makeStyles, MuiThemeProvider} from "@material-ui/core";
import React, {useEffect} from "react";
import Cookies from "js-cookie";
import {logout, roleProp} from "../login/session";
import MenuItem from "@material-ui/core/MenuItem";
import {loginTheme} from "../scenes/common";
import AppBar from "@material-ui/core/AppBar/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/core/SvgIcon/SvgIcon";
import Menu from "@material-ui/core/Menu";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Snackbar from "@material-ui/core/Snackbar";
import CloseIcon from '@material-ui/icons/Close';

const useStyles = makeStyles(theme => ({
    close: {
        padding: theme.spacing(0.5),
    },
}));

export default function Notification(props) {
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);

    useEffect(() => {
        handleClick();
    }, []);
    const handleClick = (message) => {
        setOpen(true);
    };

    const handleClose = (message) => {
        return (event, reason) => {
            if (reason === 'clickaway') {
                return;
            }

            setOpen(false);
            props.noticlosed(message);
        };
    };

    return (
        <div>
            <Snackbar
                anchorOrigin={{
                    vertical: 'bottom',
                    horizontal: 'right',
                }}
                open={open}
                autoHideDuration={6000}
                onClose={handleClose(props.message)}
                ContentProps={{
                    'aria-describedby': 'message-id',
                }}
                message={<span id="message-id">{props.message}</span>}
                action={[
                    <IconButton
                        key="close"
                        aria-label="close"
                        color="inherit"
                        className={classes.close}
                        onClick={handleClose(props.message)}
                    >
                        <CloseIcon />
                    </IconButton>,
                ]}
            />
        </div>
    );
}