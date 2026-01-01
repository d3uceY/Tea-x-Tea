<p align="center">
  <img src="./build/windows/icon.ico" alt="icon">
</p>

# TeaXTea 📝

A lightweight, cross-platform desktop text file manager built with Wails, allowing users to create, edit, view, and manage text files with a modern, beautiful interface.

## Overview

TeaXTea is a desktop application that simplifies text file management by providing an intuitive interface to work with `.txt` files in a selected directory. The application persists your chosen directory and provides a seamless editing experience with live previews and easy file operations.

## Tech Stack

### Backend
- **Go** - Backend logic and file system operations
- **Wails v2** - Go + Web frontend framework for building desktop applications
- **SQLite** (modernc.org/sqlite) - Lightweight database for storing application data (last selected directory)

### Frontend
- **React 18** - UI library for building the interface
- **Vite** - Fast build tool and development server
- **Tailwind CSS v4** - Utility-first CSS framework with modern styling
- **shadcn/ui** - Re-usable component library built on Radix UI primitives
  - Dialog, Button, Card, Input, Textarea components
- **Radix UI** - Accessible, unstyled UI primitives
- **Lucide React** - Icon library for modern icons

## Architecture

### Application Structure

```
TeaXTea/
├── main.go              # Application entry point and Wails configuration
├── app.go               # Core application logic and file operations
├── db.go                # Database initialization and table creation
├── dir_db.go            # Directory persistence logic
└── frontend/
    ├── src/
    │   ├── App.jsx      # Main React component
    │   ├── components/  # UI components (shadcn/ui)
    │   └── wailsjs/     # Generated bindings for Go ↔ React communication
    └── package.json
```

### How It Works

#### 1. **Application Initialization**
- On startup, the Go backend initializes a SQLite database in the user's config directory (`%APPDATA%/TeaXTea` on Windows)
- Creates a `last_dir_tb` table to persist the last selected directory
- Wails runtime bridges the Go backend with the React frontend

#### 2. **Directory Management**
- Users select a folder containing `.txt` files via native file dialog
- The selected directory path is saved to the SQLite database
- On subsequent launches, the app automatically loads files from the last selected directory
- Users can change the directory at any time
<img width="929" height="51" alt="image" src="https://github.com/user-attachments/assets/cd0f65c8-970e-46a0-8581-4d6bd45dc33e" />

#### 3. **File Operations**

**Read Files:**
- Scans the selected directory for `.txt` files
- Displays file list with name and size metadata
- Clicking a file loads its content into the editor
<img width="965" height="318" alt="image" src="https://github.com/user-attachments/assets/36e2e78b-c49e-44b2-ac88-804c9b6a9e29" />

**Create/Edit Files:**
- Users type content in the textarea editor
- Provide a filename in the input field
- Save triggers a native "Save File" dialog
- File is written to the selected location
<img width="1348" height="802" alt="image" src="https://github.com/user-attachments/assets/7b967455-57fb-40e2-beca-58c26070d622" />


**Delete Files:**
- Click delete icon on any file
- Confirmation dialog (shadcn Dialog component) asks for confirmation
- On confirmation, file is permanently deleted from disk
- File list refreshes automatically
<img width="1359" height="860" alt="image" src="https://github.com/user-attachments/assets/02fc71df-1c91-42da-a29d-a17f158d922b" />

#### 4. **Frontend ↔ Backend Communication**
- Wails generates JavaScript bindings from Go methods
- React components call Go functions directly: `SaveTextFile()`, `ReadFile()`, `DeleteFile()`, etc.
- No REST API needed - direct IPC (Inter-Process Communication) via Wails runtime
- Type-safe communication with generated TypeScript definitions

#### 5. **State Management**
- React `useState` hooks manage:
  - Current file content and title
  - File list from selected directory
  - Dialog open/close state
  - File to be deleted (for confirmation)
- `useEffect` hook loads directory on component mount

## Features

✨ **Modern UI** - Clean, responsive interface built with Tailwind CSS and shadcn/ui  
📁 **Directory Persistence** - Remembers your last selected folder  
📝 **In-App Editor** - Edit text files directly in the application  
💾 **Save Anywhere** - Choose where to save files with native dialogs  
🗑️ **Safe Deletion** - Confirmation dialog prevents accidental deletions  
🎨 **Cross-Platform** - Works on Windows, macOS, and Linux  
⚡ **Fast & Lightweight** - No bloat, minimal resource usage

## Development

### Prerequisites
- Go 1.24+
- Node.js 18+
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

### Run Development Server
```bash
wails dev
```

### Build for Production
```bash
wails build
```

## License

This project is a personal side project by d3uceY.
