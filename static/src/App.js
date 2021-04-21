import logo from './logo.svg';
import './App.css';

import Thing from "./components/Thing.js";

function App() {
  return (
    <div className="App">
      <Thing name="Test Thing" description="Test Thing description..." />
    </div>
  );
}

export default App;
