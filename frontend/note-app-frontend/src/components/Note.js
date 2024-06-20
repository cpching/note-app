import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import apiUrl from '../Config';
import './Note.css';

const Note = () => {
    const { id } = useParams();
    const [note, setNote] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios.get(`${apiUrl}/notes/${id}`)
            .then(response => {
                setNote(response.data);
                setLoading(false);
            })
            .catch(error => {
                console.error("There was an error fetching the note!", error);
                setLoading(false);
            });
    }, [id]);

    const saveNote = async ()=> {
        try {
            console.log("Attempting to save note:", note);
            await axios.put(`${apiUrl}/notes/${id}`, note);
            console.log("Note saved successfully");
        } catch (error) {
            console.error("Error saving the note:", error);
        }
    };

    useEffect(() => {
        const handleBeforeUnload = () => {
            if (note.title && note.content) {
                console.log("Saving note due to changes:", note);
                saveNote();
            } else {
                console.log("No changes to save.");
            }
        };

        window.addEventListener('beforeunload', handleBeforeUnload);

        return () => {
            window.removeEventListener('beforeunload', handleBeforeUnload);
        };
    }, [note]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setNote(prevNote => ({
            ...prevNote,
            [name]: value
        }));
    };

    const handleSaveClick = () => {
        saveNote().then(() => {
            // Optional navigation or state update
        });
    };

    if (loading) {
        return <p>Loading...</p>;
    }

    if (!note) {
        return <p>Note not found</p>;
    }

    return (
        /*                    <h1 className="display-4">{note.title}</h1>
                </div>
            </div>
                    <textarea
                        id="content"
                        value={note.content}
                        onChange={handleChange}
                    ></textarea>
                </div>
            </div>
        </div> */
        <div className="container mt-4">
            <div className="row">
                <div className="col title-area">
                    <input
                        className="border-0"
                        id="title"
                        type="text"
                        name="title"
                        value={note.title || ''}
                        onChange={handleChange}
                        style={{ height: 'auto', width: '100%', padding: '8px' }}
                        placeholder="Title"
                    />
                </div>
            </div>
            <div className="row">
                <div className="col content-area">
                    <textarea
                        id="content"
                        name="content"
                        className="form-control border-0 custom-textarea"
                        value={note.content || ''}
                        onChange={handleChange}
                        style={{ resize: 'none', height: 'auto', width: '100%', padding: '8px' }}
                        placeholder="Content"
                    />
                </div>
            </div>
            {/* <button onClick={handleSaveClick}>Save</button> */}
        </div>

    );
    };

export default Note;

