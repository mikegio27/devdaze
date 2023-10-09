import React, { useState, useEffect } from 'react';

const TechSkills = () => {
    const [techSkillsData, setTechSkillsData] = useState();

    useEffect(() => {
        fetch("http://localhost:8080/resume/techskills")
            .then(response => response.json())
            .then(data => {
                console.log("DATA: ", JSON.stringify(data))
                setTechSkillsData(data);
            })
            .catch(error => {
                console.error("Error fetching the techSkills data:", error);
            });
    }, []);

    if (!techSkillsData) {
        return <div>Loading Tech Skills...</div>;
    }

    return (
        <div className="resume">
            <h2>Technical Skills</h2>
            <div className="skills">
                <div>
                    <strong>Languages:</strong> {techSkillsData.Languages.join(', ')}
                </div>
                <div>
                    <strong>Platforms:</strong> {techSkillsData.Platforms.join(', ')}
                </div>
                <div>
                    <strong>Linux:</strong> {techSkillsData.Linux.join(', ')}
                </div>
                <div>
                    <strong>CI/CD:</strong> {techSkillsData.CICD.join(', ')}
                </div>
                <div>
                    <strong>Monitoring:</strong> {techSkillsData.Monitoring.join(', ')}
                </div>
                <div>
                    <strong>IAC:</strong> {techSkillsData.IAC.join(', ')}
                </div>
                <div>
                    <strong>Databases:</strong> {techSkillsData.Databases.join(', ')}
                </div>
            </div>
        </div>
    );
}

export default TechSkills;