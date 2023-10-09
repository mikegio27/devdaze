import React, { useState, useEffect } from 'react';

const Experience = () => {
    const [experienceData, setExperienceData] = useState();

    useEffect(() => {
        fetch("http://localhost:8080/resume/experience")
            .then(response => response.json())
            .then(data => {
                setExperienceData(data);
            })
            .catch(error => {
                console.error("Error fetching the experience data:", error);
            });
    }, []);

    if (!experienceData) {
        return <div>Loading Experience...</div>;
    }

    return (
        <div className="resume">
            <h2>Experience</h2>
            {experienceData.map((job, index) => (
                <div key={index} className="job">
                    <h3>{job.Title} at {job.Company}, {job.Department}</h3>
                    <p>{job.Date}</p>
                    <ul>
                        {job.Duties.map((duty, idx) => <li key={idx}>{duty}</li>)}
                    </ul>
                </div>
            ))}
        </div>
    );
}

export default Experience;