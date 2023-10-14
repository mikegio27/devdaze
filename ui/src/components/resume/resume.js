import '../../styles/App.css';
import Experience from './experience';
import Objective from './objective';
import TechSkills from './techSkills';

function Resume() {
  return (
    <div className="App">
      <header className="App-header">
        Mike Giovanetti's Resume
      </header>
        <Objective />
        <Experience />
        <TechSkills />
    </div>
    
  );
}

export default Resume;
