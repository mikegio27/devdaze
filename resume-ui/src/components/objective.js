import React, { useState, useEffect } from 'react';

const Objective = () => {
    const [objectiveData, setObjectiveData] = useState();

    useEffect(() => {
        // Replace this URL with your API endpoint
        fetch("http://34.110.219.209/resume/objective")
            .then(response => response.json())
            .then(data => {
                setObjectiveData(data.objective);
            })
            .catch(error => {
                console.error("Error fetching the objective data:", error);
            });
    }, []);

    if (!objectiveData) {
        return <div>Loading Objective...</div>;
    }

    return (
        <div className="resume">
            <h2>Objective</h2>
            <p>{objectiveData}</p>
        </div>
    );
}

export default Objective;