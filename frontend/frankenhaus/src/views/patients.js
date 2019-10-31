import React, {useEffect} from 'react';
import Paper from "@material-ui/core/Paper";
import makeStyles from "@material-ui/core/styles/makeStyles";
import "../dependencies/component"
import GenericTable from "../components/GenericTable";
import {dblist, dbremove, dbupdate} from "../api/crud";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const Patients = ({history}) => {
    const [rows, setRows] = React.useState([]);

    useEffect(() => {
        console.log("requesting data!");
        dblist("patients").then(data => setRows(data));
    }, []);

    const classes = useStyles();

    const removeAct = {
        "name": "Remove",
        "action":(item) => {
            dbremove("patients", item).then(() =>
                dblist("patients").then(data => setRows(data))
            );
        },
    };
    const addAct = {
        "name": "Add",
        "action": () => console.log("Added"),
    };
    const editAct = {
        "name": "Edit",
        "action": (item) => {
            item.editable = true;
            setActions([submitAct]);
        },
    };
    const submitAct = {
        "name": "Submit",
        "action": (item) => {
            delete item.editable;
            dbupdate("patients", item).then(() =>
                dblist("patients").then(data => setRows(data))
            );
            setActions([removeAct, editAct]);
        },
    };

    const [actions, setActions] = React.useState([removeAct, editAct]);

    const header = {
        name: "Name",
        birth_date: "Birth date",
        gender: "Gender",
        address: "Address",
        medical_record: "Medical Records",
        actions: "Actions"
    };

    return (
        <Paper className={classes.root}>
            <GenericTable className={classes.table} aria-label="simple table" header={header} data={rows}
                          actions={actions}/>
        </Paper>
    );
};

export default Patients;