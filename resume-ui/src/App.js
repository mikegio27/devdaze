import './styles/App.css';
import Home from './components/home';
import Navbar from './components/navBar';
import Swagger from './components/utilities/swagger';
import Resume from './components/resume/resume'
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
 return ( 
 <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/swagger" element={<Swagger />} />
        <Route path="/resume" element={<Resume />} />
      </Routes>
  </Router>
 )
 }
export default App;
