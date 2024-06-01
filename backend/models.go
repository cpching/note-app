package main

import (
    "time"
)

type Note struct {
    ID          int         `json:"id"`
    Title       string      `json:"title"`
    Content     string      `json:"content"`
    CreatedAt   time.Time   `json:"create_at"`
    ModifiedAt  time.Time   `json:"modified_at"`
}


