import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import apiUrl from '../Config'; // Import the apiUrl from the configuration file
// import wsUrl from '../Config'; // Import the apiUrl from the configuration file

const Note = () => {
    const { id } = useParams();
    const [note, setNote] = useState(null);
    const [loading, setLoading] = useState(true);
    const history = useNavigate();

    useEffect(() => {
        axios.get(`${apiUrl}/notes/${id}`)
            .then(response => {
                setNote(response.data);
                setLoading(false)
            })
            .catch(error => {
                console.error("There was an error fetching the note!", error);
                setLoading(false);
            })
    }, [id]);


    const saveNote = async () => {
        try {
            await axios.put(`${apiUrl}/notes/${id}`, note);
        } catch (error) {
            console.error("Error saving the note:", error);
        }
    }

    // Save the note when the component unmounts (user closes the note)
    useEffect(() => {
        const handleBeforeUnload = (event) => {
            if (note.title || note.content) { // Check if there is any content to save
                saveNote(); // Save the note
                // Uncomment the next line if you want to show a confirmation dialog
                // event.returnValue = "Are you sure you want to leave?";
            }
        };

        window.addEventListener('beforeunload', handleBeforeUnload);

        // Cleanup the event listener on component unmount
        return () => {
            window.removeEventListener('beforeunload', handleBeforeUnload);
        };
    }, [note]);

    // Handle user input to update the note's state
    const handleChange = (e) => {
        const { name, value } = e.target;
        setNote(prevNote => ({
            ...prevNote,
            [name]: value
        }));
    };

   // Optional: Save on button click
    const handleSaveClick = () => {
        saveNote().then(() => {
            // history.push('/notes'); // Navigate to the notes list after saving
        });
    };

    if(loading) {
        return <p>Loading...</p>;
    }

    if(!note) {
        return <p>Note not found</p>;
    }

    return (
        <div>
        <h1>Edit Note</h1>
        <input
        type="text"
        name="title"
        value={note.title}
        onChange={handleChange}
        placeholder="Title"
        />
        <textarea
        name="content"
        value={note.content}
        onChange={handleChange}
        placeholder="Content"
        />
            <button onClick={handleSaveClick}>Save</button>
        </div> 
    );
};

export default Note;

