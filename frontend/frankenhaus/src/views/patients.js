import React from 'react';
import Paper from "@material-ui/core/Paper";
import makeStyles from "@material-ui/core/styles/makeStyles";
import "../dependencies/component"
import GenericTable from "../components/GenericTable";

function createData(name, calories, fat, carbs, protein) {
    return {name, calories, fat, carbs, protein};
}

const rows = [
    createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
    createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
    createData('Eclair', 262, 16.0, 24, 6.0),
    createData( 'Cupcake', 305, 3.7, 67, 4.3),
    createData( 'Gingerbread', 356, 16.0, 49, 3.9),
];

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

    return (
        <Paper className={classes.root}>
            <GenericTable className={classes.table} aria-label="simple table" data={rows} actions={actions}/>
        </Paper>
    );
};

export default Patients;