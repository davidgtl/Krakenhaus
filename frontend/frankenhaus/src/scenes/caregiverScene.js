import makeStyles from "@material-ui/core/styles/makeStyles";
import Patients from "../views/patients";
import Medications from "../views/medications";
import Caregivers from "../views/caregivers";
import Cookies from "js-cookie";
import React, {useEffect, useState} from "react";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import Notification from "../components/Notification";

const useStyles = makeStyles({
    root: {
        width: '100%',
        overflowX: 'auto',
    },
    table: {
        minWidth: 650,
    },
});


const client = new W3CWebSocket('ws://localhost:8080/ws');


const CaregiverScene = ({history, props}) => {
    const classes = useStyles();

    const [notifications, setNotifications] = useState(["Patient #1 has been sleeping for too much!"]);

    useEffect(() =>{
        client.onopen = () => {
            console.log('WebSocket Client Connected');
        };
        client.onmessage = (message) => {
            console.log(message);
            setNotifications([...notifications, message.data])
        };
    }, []);

    const handleNotiClosed = (message) => {
        setNotifications(notifications.filter(x => x != message))
    };
    const viewMap = {
        "": <Patients/>,
        "patients": <Patients/>,
        "medications": <Medications/>,
        "caregivers": <Caregivers/>,
    };
    let viewDOM = "";
    let view = Cookies.get("view");
    if (view === undefined || view == "patients")
        viewDOM = <div>{notifications.map(elem => <Notification message={elem} noticlosed={handleNotiClosed}/>)}<Patients/></div>;


    return viewDOM;
};

export default CaregiverScene;