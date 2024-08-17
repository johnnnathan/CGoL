# Conway's Game of Life - Go Server

## Overview

This project implements Conway's Game of Life in Go, using a web server to serve the game state as an HTML table. The server supports dynamic updates to the game state and visualizes the board using HTML and CSS.

## Features

- **Game of Life Simulation**: Simulates Conway's Game of Life with a configurable board size.
- **Dynamic Updates**: Continuously updates the game state and reflects changes on the web page.
- **Parallel Processing**: Utilizes concurrent processing to handle game state updates efficiently.
- **Web Interface**: Displays the game board as an HTML table, with live updates.

## Requirements

- Go 1.18 or later
- Web browser for viewing the game board

## Setup

1. **Clone the Repository**

    ```bash
    git clone https://github.com/yourusername/game-of-life.git
    cd game-of-life
    ```

2. **Build and Run**

    Build and run the Go server:

    ```bash
    go build -o game-of-life
    ./game-of-life
    ```

3. **Access the Game**

    Open your web browser and navigate to [http://localhost:8080](http://localhost:8080) to view and interact with the game board.

## Code Structure

- `main.go`: Contains the core logic of the game, including the server setup, game simulation, and HTTP handlers.
- `static/index.html`: The HTML file served by the server.

## Game Mechanics

- **Initialization**: The board is initialized with random values of 0 and 1.
- **Update**: The board updates based on Conway's Game of Life rules, with parallel processing to handle large boards.
- **Display**: The board is rendered as an HTML table, with live updates to reflect changes in the game state.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
