import axios from 'axios';
import {axiosConfig, login} from "../login/session";

const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.target);

    fetch('/api/form-submit-url', {
        method: 'POST',
        body: data,
    });
};

const dblist = (routeName) => new Promise((resolve, reject) => {
    axios.get(`http://127.0.0.1:8080/${routeName}/`, axiosConfig())
        .then(function (response) {
            resolve(response.data.object)
        })
        .catch(function (error) {
            console.log(error);
        });
});

const dbupdate = (routeName, data) => new Promise((resolve, reject) => {
    axios.post(`http://127.0.0.1:8080/${routeName}/update`, data, axiosConfig())
        .then(function (response) {
            resolve(response.data)
        })
        .catch(function (error) {
            console.log(error);
        });
});

const dbremove = (routeName, data) => new Promise((resolve, reject) => {
    axios.post(`http://127.0.0.1:8080/${routeName}/delete`, data, axiosConfig())
        .then(function (response) {
            resolve(response.data)
        })
        .catch(function (error) {
            console.log(error);
        });
});



export {
    dblist,
    dbupdate,
    dbremove
}