import './App.css';
import Home from './components/home';
import Navbar from './components/navBar';
import Swagger from './components/utilities/swagger';
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
 return ( 
 <Router>
    <div className="App">
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/swagger" element={<Swagger />} />
      </Routes>
    </div>
  </Router>
 )
 }
export default App;
