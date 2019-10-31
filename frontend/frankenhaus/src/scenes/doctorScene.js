import Paper from "@material-ui/core/Paper";
import GenericTable from "../components/GenericTable";
import React from "react";
import makeStyles from "@material-ui/core/styles/makeStyles";
import {loginTheme} from "./common";
import Button from "@material-ui/core/Button";
import {MuiThemeProvider} from "@material-ui/core";
import Patients from "../views/patients";
import Medications from "../views/medications";
import Caregivers from "../views/caregivers";
import Cookies from 'js-cookie'

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const DoctorScene = ({history, props}) => {
    const classes = useStyles();

    const viewMap = {
        "": <Patients/>,
        "patients": <Patients/>,
        "medications": <Medications/>,
        "caregivers": <Caregivers/>,
    };
    let viewDOM = "";
    let view = Cookies.get("view");
    if (view === undefined || view == "patients")
        viewDOM = <Patients/>;
    if (view === "medications")
        viewDOM = <Medications/>;
    if (view === "caregivers")
        viewDOM = <Caregivers/>;


    return viewDOM;
};

export default DoctorScene;