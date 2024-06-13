import React, { useState, useEffect } from 'react';
import axios from 'axios';
import FloatingButton from './FloatingButton';
import FloatingEditPage from './FloatingEditPage';
import apiUrl from '../Config'; // Import the apiUrl from the configuration file

const NoteList = () => {
    const [notes, setNotes] = useState([]);
    const [editingNote, setEditingNote] = useState(null);
    const [isEditing, setIsEditing] = useState(false)

    useEffect(() => {
        fetchNotes();
    }, []);

    const openEditPage = () => {
        setEditingNote(null)
        setIsEditing(true)
    };

    const closeEditPage = () => {
        setIsEditing(false);
    };

    const fetchNotes = async () => {
        const response = await axios.get(`${apiUrl}/notes`);
        setNotes(response.data);
    };

    const handleDelete = async (id) => {
        await axios.delete(`/api/notes/${id}`);
        fetchNotes();
    };

    const handleEdit = (note) => {
        setEditingNote(note);
    };

    const handleSave = () => {
        setEditingNote(null);
        fetchNotes();
    };

    return (
        <div>
            <h1>All Notes</h1>
            <FloatingButton onClick={openEditPage} />
            <FloatingEditPage 
            isVisible={isEditing} 
            onClose={closeEditPage}
            />
        </div>
    );
};

export default NoteList;

