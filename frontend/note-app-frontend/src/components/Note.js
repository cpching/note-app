import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';

const Note = () => {
    const { id } = useParams();
    const [note, setNote] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios.get(`http://localhost:8080/notes/${id}`)
            .then(response => {
                setNote(response.data);
                setLoading(false)
            })
            .catch(error => {
                console.error("There was an error fetching the note!", error);
                setLoading(false);
            })
        
    }, [id]);

    if(loading) {
        return <p>Loading...</p>;
    }

    if(!note) {
        return <p>Note not found</p>;
    }

    return (
        <div>
            <h1>{note.title}</h1>
            <p>{note.content}</p>
        </div>

    );
};

export default Note;

