import '../styles/App.css';

function Home() {
  return (
    <div className="App">
      <header className="App-header">
        Welcome to DevDaze
      </header>
        <p>
          The purpose of this application is for me to continue to sharpen my skills with web development and infrastructure.
          This application is hosted on a GKE cluster with a container for the UI, API, and noSQL database.
          It is an overengineered way for me to host my resume. I plan on continuing to add things to it to improve it and
          continue to gain experience with various aspects of full stack development and DevOps.  
        </p>
    </div>
    
  );
}

export default Home;
