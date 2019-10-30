import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Table from "@material-ui/core/Table";
import React from "react";
import Button from "@material-ui/core/Button";
import ButtonGroup from "@material-ui/core/ButtonGroup";

export default function GenericTable(props) {
    const [data] = React.useState(props.data);

    const dataFields = Object.getOwnPropertyNames(data[0]);
    const actionFields = props.actions.length > 0 ? ["Actions"] : [];
    const header = dataFields.concat(actionFields).map(x => <TableCell align="center">{x}</TableCell>);
    const body = data.map(elem => (
        <TableRow>
            {Object.getOwnPropertyNames(elem).map(key => <TableCell align="center">{elem[key]}</TableCell>)}
            {
                <TableCell align="center">
                    <ButtonGroup color="primary" variant="contained" aria-label="contained secondary button group">
                        {
                            props.actions.map(action => <Button onClick={() => action.action()}>{action.name}</Button>)
                        }
                    </ButtonGroup>
                </TableCell>}
        </TableRow>
    ));

    return (
        <Table aria-label="simple table">
            <TableHead>
                <TableRow>
                    {header}
                </TableRow>
            </TableHead>
            <TableBody>
                {body}
            </TableBody>
        </Table>
    );
}