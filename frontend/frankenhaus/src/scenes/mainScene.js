import React, {useState} from "react";
import {createMuiTheme, MuiThemeProvider} from "@material-ui/core";
import CssBaseline from "@material-ui/core/CssBaseline";
import {roleProp} from "../login/session";
import {Route, Link, BrowserRouter as Router} from 'react-router-dom'
import Container from "@material-ui/core/Container";

const routing = (view) => {

    return (
        <Router>
            <Route exact path="/" component={roleProp("homeScene")(view)}/>
        </Router>
    );

};

function MainScene() {
    const [currentView, setCurrentView] = useState("");

    return (
        <Container>
            {roleProp("navbar")(setCurrentView)}
            {routing(currentView)}
        </Container>)
}

export {
    MainScene
}