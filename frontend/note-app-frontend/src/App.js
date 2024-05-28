import React from 'react';
import './App.css';
import NoteList from './components/NoteList';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Note App</h1>
      </header>
      <NoteList />
    </div>
  );
}

export default App;

