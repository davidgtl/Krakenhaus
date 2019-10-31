/* eslint-disable no-unused-vars */
import React from 'react';
import ReactDOM from 'react-dom';
import {Route, Link, BrowserRouter as Router} from 'react-router-dom'
import {createMuiTheme} from '@material-ui/core/styles';
import './index.css';
import App from './App';
import Caregiver from './views/caregiver'
import Patients from './views/patients'

import * as serviceWorker from './serviceWorker';
import {ThemeProvider} from "@material-ui/styles";
import Paper from "@material-ui/core/Paper";
import Container from "@material-ui/core/Container";
import {MuiThemeProvider} from "@material-ui/core";
import CssBaseline from "@material-ui/core/CssBaseline";
import Login from "./login/login";
import Medications from "./views/medications";
import DoctorScene from "./scenes/doctorScene";
import CaregiverScene from "./scenes/caregiverScene";
import PatientScene from "./scenes/patientScene";
import Navbar from "./components/navbar";
import {roleProp} from "./login/session";

const routing = () => {

    return (
        <Router>
            <Route exact path="/" component={roleProp("homeScene")}/>
        </Router>
    );
};

const theme = createMuiTheme({
    palette: {
        type: 'dark',
        primary: {main: '#26b0f5'},
        secondary: {main: '#4DD0E1'}
    },
    typography: {
        fontFamily: [
            'Roboto',
            '-apple-system',
            'BlinkMacSystemFont',
            '"Segoe UI"',
            '"Helvetica Neue"',
            'Arial',
            'sans-serif',
            '"Apple Color Emoji"',
            '"Segoe UI Emoji"',
            '"Segoe UI Symbol"',
        ].join(',')
    },
    background: {
        default: '#181c26'
    }
});

ReactDOM.render(
    <MuiThemeProvider theme={theme}>
        <CssBaseline/>
        {roleProp("navbar")}

        {routing()}
    </MuiThemeProvider>, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
