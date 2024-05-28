import React, { useState, useEffect } from 'react';
import axios from 'axios';
import NoteForm from './NoteForm';

const NoteList = () => {
  const [notes, setNotes] = useState([]);
  const [editingNote, setEditingNote] = useState(null);

  useEffect(() => {
    fetchNotes();
  }, []);

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
      <h1>Notes</h1>
      {editingNote ? (
        <NoteForm note={editingNote} onSave={handleSave} />
      ) : (
        <NoteForm onSave={handleSave} />
      )}
      <ul>
        {notes.map((note) => (
          <li key={note.id}>
            <h2>{note.title}</h2>
            <p>{note.content}</p>
            <button onClick={() => handleEdit(note)}>Edit</button>
            <button onClick={() => handleDelete(note.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default NoteList;

