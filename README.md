# TCC

Interview project based on `No9.docx` using:

- Go for the backend API
- Angular for the frontend UI
- A lightweight file-backed JSON store for comment persistence

## Structure

- `backend`: Go HTTP API
- `frontend`: Angular application

## Run

### Backend

```bash
cd backend
go run .
```

The API starts on `http://localhost:8080`.

### Frontend

```bash
cd frontend
npm install
npm start
```

The app starts on `http://localhost:4200`.

## Behavior

- The page reproduces the reference layout from the document.
- The post is shown with the original image and metadata.
- When `Blend 285` types a comment and presses `Enter`, the new text appears immediately below the input.
- Submitted comments are also persisted by the backend in `backend/data/comments.json`.
