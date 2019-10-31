import React, {useEffect} from 'react';
import Paper from "@material-ui/core/Paper";
import makeStyles from "@material-ui/core/styles/makeStyles";
import "../dependencies/component"
import GenericTable from "../components/GenericTable";
import {dblist, dbremove, dbupdate} from "../api/crud";
import Button from "@material-ui/core/Button";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
    newentry:{
        float: 'right',
        margin: '1em',
    }
});


const Caregivers = ({history}) => {
    const apiroute = "caregivers";
    const [rows, setRows] = React.useState([]);

    useEffect(() => {
        dblist(apiroute).then(data => setRows(data));
    }, []);

    const classes = useStyles();

    const removeAct = {
        "name": "Remove",
        "action":(item) => {
            dbremove(apiroute, item).then(() =>
                dblist(apiroute).then(data => setRows(data))
            );
        },
    };
    const editAct = {
        "name": "Edit",
        "action": (item) => {
            item.actionstate = "editable";
        },
    };
    const cancelAct = {
        "name": "Cancel",
        "action": (item) => {
            delete item.actionstate;
            dblist(apiroute).then(data => setRows(data));
        },
    };
    const submitAct = {
        "name": "Submit",
        "action": (item) => {
            delete item.actionstate;
            dbupdate(apiroute, item).then(() =>
                dblist(apiroute).then(data => setRows(data))
            );
        },
    };

    const addNewEntry = () => {
        const shell = {
            "name" : "",
            "birth_date" : "2012-01-01T23:28:56.782Z",
            "gender": "",
            "address": "",
            actionstate: "editable"
        };
        setRows([...rows, shell]);
    };

    const actions ={
        "": [editAct, removeAct],
        "editable": [submitAct, cancelAct]
    };

    const header = {
        name: "Name",
        birth_date: "Birth date",
        gender: "Gender",
        address: "Address",
        actions: "Actions"
    };

    return (
        <Paper className={classes.root}>
            <Button className={classes.newentry} variant="contained" color="secondary" onClick={() =>  addNewEntry()}>
                New Entry
            </Button>
            <GenericTable className={classes.table} aria-label="simple table" header={header} data={rows}
                          actions={actions}/>
        </Paper>
    );
};

export default Caregivers;