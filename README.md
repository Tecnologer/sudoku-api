# Sudoku API

Expose Sudoku game as ResAPI.

## Dependencies

- [sudoku][1]: `go get -u github.com/tecnologer/sudoku-lib`

## Run

### From Binary

- [Download][2] the version for your system
- `cd <download_path>`
- run `./<os>-sudoku-api-<arch>-<version>`
  - I.e. `./linux-sudoku-api-amd64-0.0.1`

### From Source Code

- `go get -u github.com/tecnologer/sudoku-api`
- `cd $GOPATH/src/github.com/tecnologer/sudoku-api`
- `go build` or `make`
- `./sudoku-api`

### Flags

- `-port` **(int)**: Starting port server (default 8088)
- `-v` **(bool)**: Enanble verbouse log
- `-version` **(bool)**: Returns the version of the build

```bash

./sudoku-api -version
# output
> 0.0.1.210210

# run verbouse in port 80
./sudoku-api -v -port 80

# output
INFO[0000] Starting server on :80...
.
.
.
```

## Endpoints

- `/api/game?level=<level>`: returns a game with the level specified

  - Response:

  ```json
  {
    "board": "number[][]",
    "level": "<string>",
    "start_time": "<datetime>",
    "locked_coordinates": "[{ x, y }]"
  }
  ```

- `/api/game/set?x=<x>&y=<y>&n=<cell_value>`: sets the cell value in the coordinate

  - Response:

  ```json
  {
    "board": "number[][]",
    "level": "<string>",
    "start_time": "<datetime>",
    "locked_coordinates": "[{x,y}]"
  }
  ```

- `/api/game/levels`: returns the list of available levels

  - Response:

  ```json
    string[]
  ```

- `/api/game/validate[?empty]`: validates if the board is complete correctly

  - Response: _an object with the total of errors and a dictionary specifing the type and coordinate of the error_

  ```json
  {
    "errors": {
      "square": "[{x,y}]",
      "row": "[{x,y}]",
      "column": "[{x,y}]",
      "empty": "[{x,y}]"
    },
    "count": "number"
  }
  ```

[1]: https://github.com/Tecnologer/sudoku-lib
[2]: https://github.com/Tecnologer/sudoku-api/releases
