# Ascii-Art-Web

## Overview

Ascii-Art-Web is a web-based version of the Ascii-Art project built with Go. It allows users to enter text through a graphical user interface (GUI), select an ASCII art banner style, and generate ASCII art directly in the browser.

The application runs a Go HTTP server that handles requests and renders ASCII art using different banner templates.

---

## Features

* Web-based graphical user interface
* Generate ASCII art from user input
* Support for multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* HTTP server built with Go
* Dynamic HTML rendering using Go templates
* Form submission with POST requests

---

## HTTP Endpoints

### GET /

Displays the main webpage.

The homepage contains:

* A text input field
* Banner selection options (radio buttons, select menu, etc.)
* A submit button


### POST /ascii-art

Receives:

* User text
* Selected banner

Processes the input and generates the corresponding ASCII art before displaying the result.

---

## Project Structure

```text
ascii-art-web/
│
├── main.go
├── templates/
|   |__result.html
│   └── index.html
│
│──Ascii.go
├──Handler.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
│
└── README.md
```

---

## Requirements

* Go 1.22.2 

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd ascii-art-web
```

Run the server:

```bash
go run .
```

---

## Usage

1. Start the server.
2. Open your browser.
3. Visit:

```text
http://localhost:8080
```

4. Enter text in the input field.
5. Select a banner style:

   * Standard
   * Shadow
   * Thinkertoy
6. Click the Generate button.
7. View the generated ASCII art.

---

## Example

Input:

```text
Hello
```

Banner:

```text
standard
```

Output:

```text
 _   _      _ _
| | | | ___| | | ___
| |_| |/ _ \ | |/ _ \
|  _  |  __/ | | (_) |
|_| |_|\___|_|_|\___/
```

---

## Technologies Used

* Go
* HTML
* Go Templates
* HTTP Server

---

## Learning Objectives

This project demonstrates:

* HTTP server creation in Go
* Handling GET and POST requests
* HTML form processing
* Template rendering
* User input validation
* Working with ASCII art generators
* Web application development fundamentals

---

## Authors

* Kpensalen Aondongu
* Benjamin Otete
* Odoh Emmanuel