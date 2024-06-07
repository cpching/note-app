import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { BrowserRouter } from 'react-router-dom'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.css'
import NoteList from './components/NoteList';
import Note from './components/Note';
import Home from './components/Home';

const root = ReactDOM.createRoot(document.getElementById('root'));

root.render( <App />)


// 建立 Router （用 createBrowserRouter 建立 BrowserRouter)


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
