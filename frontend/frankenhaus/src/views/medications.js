import React from 'react';

const Medications = () => {
    const [value, setValue] = React.useState('');
    const onChange = event => setValue(event.target.value);
    return (
        <div>
            <h1>Medication list</h1>
            <input value={value} type="text" onChange={onChange} />
            <p>{value}</p>
        </div>
    );
};

export default Medications;