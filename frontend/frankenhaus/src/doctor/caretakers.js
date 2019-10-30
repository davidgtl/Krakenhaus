import React from 'react';

const Caretakers = () => {
    const [value, setValue] = React.useState('');
    const onChange = event => setValue(event.target.value);
    return (
        <div>
            <h1>Caretaker list</h1>
            <input value={value} type="text" onChange={onChange} />
            <p>{value}</p>
        </div>
    );
};

export default Caretakers;