const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.target);

    fetch('/api/form-submit-url', {
        method: 'POST',
        body: data,
    });
};

const list = (routeName) => {

};


export {
    list
}