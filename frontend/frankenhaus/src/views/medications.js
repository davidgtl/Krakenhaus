import React, {useEffect} from 'react';
import Paper from "@material-ui/core/Paper";
import makeStyles from "@material-ui/core/styles/makeStyles";
import "../dependencies/component"
import GenericTable from "../components/GenericTable";
import {dblist} from "../api/crud";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const Medications = ({history}) => {
    const [rows, setRows] = React.useState([]);

    useEffect(() => {
        dblist("caregivers").then(data => setRows(data));
    });

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
                    "action": console.log
                }
            ]
        }
    ];

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
            <GenericTable className={classes.table} aria-label="simple table" header={header} data={rows} actions={actions}/>
        </Paper>
    );
};

export default Medications;