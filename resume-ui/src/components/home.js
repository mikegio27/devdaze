import '../App.css';
import Experience from './resume/experience';
import Objective from './resume/objective';
import TechSkills from './resume/techSkills';

function Home() {
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

export default Home;
