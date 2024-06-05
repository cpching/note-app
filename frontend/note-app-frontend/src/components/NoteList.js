import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Note from './Note';
import FloatingButton from './FloatingButton';
import FloatingEditPage from './FloatingEditPage';

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
        const response = await axios.get('http://localhost:8080/notes');
        setNotes(response.data);
    };

    const handleDelete = async (id) => {
        await axios.delete(`http://localhost:8080/notes/${id}`);
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
        {notes.map(note => (
            <Note noteId = {note.id} noteTitle = {note.title} noteContent = {note.content} />
        ))}
        <FloatingButton onClick={openEditPage} />
        <FloatingEditPage 
        isVisible={isEditing} 
        onClose={closeEditPage}
        />
        </div>

    );
};


export default NoteList;

