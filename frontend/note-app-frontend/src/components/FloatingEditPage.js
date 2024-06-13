import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './FloatingEditPage.css';
import apiUrl from '../Config';

const FloatingEditPage = ({ isVisible, onClose }) => {
    const handleCreate = async (event) => {
        event.preventDefault();
        await axios.post(`${apiUrl}/notes`, { title, content })
    } 

    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    
    if (!isVisible) return null;

    return (
        <div className="floating-edit-page">
            <div className="edit-container">
                <input className='edit-title'
                  placeholder="Title"
                value={title} 
                onChange={(e) => setTitle(e.target.value)}
                />
                <textarea className='edit-content'
                  placeholder="Content"
                value={content} 
                onChange={(e) => setContent(e.target.value)}
                />
        
                <div className="edit-buttons">
                    <button onClick={handleCreate}>Save</button>
                    <button onClick={onClose}>Cancel</button>
                </div>
            </div>
        </div> 
    );
}

export default FloatingEditPage

