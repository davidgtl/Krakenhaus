import Cookies from 'js-cookie'
import Navbar from "../components/navbar";
import DoctorScene from "../scenes/doctorScene";
import CaregiverScene from "../scenes/caregiverScene";
import Login from "./login";
import React from "react";

export const axiosConfig = () => {
    const jwt = Cookies.get('__session');
    return {
        headers: {
            Authorization: "Bearer " + Cookies.get("token")
        }
    };
};

export const login = (user) => {
    console.log(user);
    console.log(user.token);
    Cookies.set("token", user.token);
    Cookies.set("role", user.role);
    window.location.reload();
};

export const logout = () => {
    Cookies.remove("token");
    Cookies.remove("role");
    window.location.reload();
};

export const getSession = () => {
    const jwt = Cookies.get('__session');
    let session;
    try {
        if (jwt) {
            const base64Url = jwt.split('.')[1];
            const base64 = base64Url.replace('-', '+').replace('_', '/');
            session = JSON.parse(window.atob(base64))
        }
    } catch (error) {
        console.log(error)
    }
    return session
};


const roleProperties = () => {
    return {
        default: {
            navbar: <Navbar />,
        },
        "doctor": {
            homeScene: DoctorScene,
        },
        "caregiver": {
            homeScene: CaregiverScene,
        },
        "patient": {
            homeScene: CaregiverScene
        },
        "": {
            homeScene: Login,
            navbar: () => {}
        }
    };
};

//const getRole = () => Cookie.get("role");

export const roleProp = (prop) =>{
    let role = Cookies.get("role");
    if(role === undefined)  role = "";

    const properties = roleProperties();
    console.log(properties);
    if(properties[role][prop] === undefined)
        return properties.default[prop];

    return properties[role][prop];
};
