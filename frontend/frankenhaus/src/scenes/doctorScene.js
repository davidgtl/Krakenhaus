import Paper from "@material-ui/core/Paper";
import GenericTable from "../components/GenericTable";
import React from "react";
import makeStyles from "@material-ui/core/styles/makeStyles";
import {loginTheme} from "./common";
import Button from "@material-ui/core/Button";
import {MuiThemeProvider} from "@material-ui/core";
import Patients from "../views/patients";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const DoctorScene = ({history}) => {
    const classes = useStyles();
    const actions = [
        {
            "name": "Remove",
            "action": () => console.log("Removed")
        },
        {
            "name": "Add",
            "action": () => console.log("Added")
        },
        {
            "name": "Edit",
            "action": () =>  history.push('/'),
            "subactions": [
                {
                    "name": "Save",
                    "action":  console.log
                }
            ]
        }
    ];

    return (
        <Patients/>
    );
};

export default DoctorScene;