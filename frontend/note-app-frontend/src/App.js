import React from 'react';
// import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import './App.css';
import NoteList from './components/NoteList';
import Note from './components/Note';
import Home from './components/Home';

const router = createBrowserRouter([
  { path: "/", element: <Home/> },
  {
    path: "/notes/",
    element: <NoteList />,
  },
  {
      path: "/notes/:id",
    element: <Note />,
  },
]);

function App() {
  return (
    <div>
      <RouterProvider router={router} />
    </div>

  );
}

export default App;

