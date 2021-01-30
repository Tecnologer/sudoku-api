# Sudoku API

Expose Sudoku game as ResAPI.

## Dependencies

- [sudoku][1]: `go get -u github.com/tecnologer/sudoku-lib`

## Endpoints

- `/api/game?level=<level>`: returns a game with the level specified

  - Response:

  ```json
    {
        "board": number[][],
        "level": "<string>",
        "start_time": "<datetime>",
        "locked_coordinates": [{x,y}]
    }
  ```

- `/api/game/set?x=<x>&y=<y>&n=<cell_value>`: sets the cell value in the coordinate

  - Response:

  ```json
    {
        "board": number[][],
        "level": "<string>",
        "start_time": "<datetime>",
        "locked_coordinates": [{x,y}]
    }
  ```

- `/api/game/levels`: returns the list of available levels

  - Response:

  ```json
    string[]
  ```

- `/api/game/validate[?empty]`: validates if the board is complete correctly

  - Response: _an object with the total of errors and a dictionary specifing the type and coordinate of the error_

  ```json{
    "errors": {
        "square": [{x,y}],
        "row": [{x,y}],
        "column": [{x,y}],
        "empty": [{x,y}],
    },
    "count": number
  }
  ```

[1]: https://github.com/Tecnologer/sudoku-lib
