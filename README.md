# robovsdino
The following information has been given by Grover for the coding challenge:

## These are the features required:

- Be able to create an empty simulation space - an empty 50 x 50 grid;
- Be able to create a robot in a certain position and facing direction;
- Be able to create a dinosaur in a certain position;
- Issue instructions to a robot - a robot can turn left, turn right, move forward, move backward, and attack;
- A robot attack destroys dinosaurs around it (in front, to the left, to the right or behind);
- No need to worry about the dinosaurs - dinosaurs don't move;
- Display the simulation's current state;
- Two or more entities (robots or dinosaurs) cannot occupy the same position;
- Attempting to move a robot outside the simulation space is an invalid operation.

## Things we are looking for

- Immutability/Referential transparency;
- Idiomatic code;
- Adherence to community/standard library style guides;
- Separation of concerns;
- Unit and integration tests;
- API design;
- Domain modeling;
- Attention to possible concurrency issues;
- Error handling.

## Installation
1. This project needs Go or Docker installed
2. I haven't deployed to the public because it is a challenge, it may be inappropriate

### Local machine run
1. If you haven't installed go, please install: https://golang.org/doc/install
2. You can run the tests by: `go test` or more verbose: `go test -v`
3. You need setup 2 environment variables: `PORT=8080` (or your taste) and `APIENV=PROD` (or `APIENV=DEV`)
4. You can run the app by: `go run .` or you can build then run if you want: `go build -o robovsdino` and `./robovsdino`
5. When app runs, you can see possible API calls in the terminal if `APIENV=DEV`
6. For more info about code and how to use it, please refer its documentation

docker build -t canartuc/robovsdino .

docker run -p 8080:8080 canartuc/robovsdino

