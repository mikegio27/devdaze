import './App.css';
import Experience from './components/experience';
import Objective from './components/objective';
import TechSkills from './components/techSkills';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        Mike Giovanetti's Resume
      </header>
      <div>
        <Objective />
      </div>
        <Experience />
      <div>
      </div>
        <TechSkills />
      <div>
      </div>
    </div>
    
  );
}

export default App;
