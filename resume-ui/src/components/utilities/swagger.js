import React, { useState, useEffect } from 'react';

const Swagger = () => {
    const [swaggerData, setSwaggerData] = useState();

    useEffect(() => {
        // Replace this URL with your API endpoint
        fetch("https://api.devdaze.org/swagger/index.html")
            .then(response => response.json())
            .then(data => {
                setSwaggerData(data);
            })
            .catch(error => {
                console.error("Error fetching the objective data:", error);
            });
    }, []);

    if (!swaggerData) {
        return <div>Loading swagger...</div>;
    }

    return (
        <div>
            {swaggerData}
        </div>
    );
}

export default Swagger;