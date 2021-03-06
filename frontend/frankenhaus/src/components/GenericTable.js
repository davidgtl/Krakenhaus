import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Table from "@material-ui/core/Table";
import React, {useState} from "react";
import Button from "@material-ui/core/Button";
import ButtonGroup from "@material-ui/core/ButtonGroup";
import {makeStyles} from "@material-ui/core";
import Input from "@material-ui/core/Input";
import {logout} from "../login/session";
import Toolbar from "@material-ui/core/Toolbar";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles(theme => ({
    editable: {
        display: 'block',
    },
    hidden: {
        display: 'none',
    },
}));

export default function GenericTable(props) {
    const classes = useStyles();

    const [refreshState, setRefresh] = useState(false);

    let data = props.data;

    const isEditable = (x) => {
        if (x === undefined)
            return false;
        if (x.actionstate === undefined)
            return false;
        return x.actionstate === "editable";
    };

    const handleChange = (event, elem, key) => {
        elem[key] = event.target.value;
    };

    const headerNames = Object.getOwnPropertyNames(props.header);
    const headerDOM = headerNames.map(x => <TableCell align="center">{props.header[x]}</TableCell>);

    const createBody = () => data.map(elem => {
        const editableClass = isEditable(elem) ? classes.editable : classes.hidden;
        const noneditableClass = !isEditable(elem) ? classes.editable : classes.hidden;
        let cols = Object.getOwnPropertyNames(elem).filter(x => headerNames.includes(x)).map(key =>
            <TableCell refreshstate={"" + refreshState} align="center">
                <span className={noneditableClass}>{elem[key]}</span>
                <Input
                    onChange={event => handleChange(event, elem, key)}
                    defaultValue={elem[key]}
                    className={editableClass}
                    inputProps={{
                        'aria-label': 'description',
                    }}/>
            </TableCell>);

        let actions = props.actions[""];

        if (elem.actionstate !== undefined)
            actions = props.actions[elem.actionstate];


        return (
            <TableRow>
                {cols}
                {<TableCell align="center">
                    <ButtonGroup color="primary" variant="contained" aria-label="contained secondary button group">
                        {
                            actions.map(action => <Button
                                onClick={() => {
                                    action.action(elem);
                                    setRefresh(!refreshState)
                                }}>{action.name}</Button>)
                        }
                    </ButtonGroup>
                </TableCell>}
            </TableRow>
        );
    });

    return (
        <Table stickyHeader aria-label="simple table">
            <TableHead>
                <TableRow>
                    {headerDOM}
                </TableRow>
            </TableHead>
            <TableBody>
                {createBody()}
            </TableBody>
        </Table>
    );
}