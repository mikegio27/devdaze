import React from 'react';
import { Link } from 'react-router-dom';

function Navbar() {
    return (
      <div className="navbar">
        <Link to="/">Home </Link>
        <Link to="/resume">Resume</Link>
        <Link to="/swagger">Swagger</Link>
        <a href="https://github.com/mikegio27/resume-app" target="_blank" rel="noopener noreferrer">GitHub</a>
      </div>
    );
  }
export default Navbar;