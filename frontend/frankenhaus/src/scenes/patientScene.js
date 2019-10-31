import makeStyles from "@material-ui/core/styles/makeStyles";
import Patients from "../views/patients";
import Medications from "../views/medications";
import Caregivers from "../views/caregivers";
import Cookies from "js-cookie";
import React from "react";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const PatientsScene = ({history, props}) => {
    const classes = useStyles();

    const viewMap = {
        "": <Medications/>,
        "medications": <Medications/>,
    };
    let viewDOM = "";
    let view = Cookies.get("view");
    if (view === undefined || view == "patients")
        viewDOM = <Patients/>;


    return viewDOM;
};

export default PatientsScene;