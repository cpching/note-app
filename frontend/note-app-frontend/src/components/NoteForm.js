import React, { useState, useEffect } from 'react';
import axios from 'axios';

const NoteForm = ({ note, onSave }) => {
  const [title, setTitle] = useState(note ? note.title : '');
  const [content, setContent] = useState(note ? note.content : '');

  const handleSubmit = async (event) => {
    event.preventDefault();
    if (note) {
      await axios.put(`http://localhost:8080/notes/${note.id}`, { title, content });
    } else {
      await axios.post('http://localhost:8080/notes', { title, content });
    }
    onSave();
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Title</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      </div>
      <div>
        <label>Content</label>
        <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
      </div>
      <button type="submit">Save</button>
    </form>
  );
};

export default NoteForm;

