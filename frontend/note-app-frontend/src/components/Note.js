import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Note = ({noteId, noteTitle, noteContent}) => {
    return (
        <h1>{noteTitle}</h1>

    );

};

export default Note

