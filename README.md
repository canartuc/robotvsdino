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
3. Personally, I tested in Ubuntu 20.10 and macOS Catalina 10.15.7 but not tested in Windows so please feel free to 
find your way

### Local machine run
1. If you haven't installed go, please install: https://golang.org/doc/install
2. Run the command inside the project: `go mod download`
3. You can run the tests by: `go test` or more verbose: `go test -v`
4. You need setup 2 environment variables: `PORT=8080` (or your taste) and `APIENV=PROD` (or `APIENV=DEV`)
5. You can run the app by: `go run .` or you can build then run if you want: `go build -o robovsdino` and `./robovsdino`
6. When app runs, you can see possible API calls in the terminal if `APIENV=DEV`
7. For more info about code and how to use it, please refer its documentation **godoc.txt**. For API calls, please check here.

### Docker run
1. If you haven't installed docker, please install: https://docs.docker.com/get-docker/
2. Inside the project folder, run the command: `docker build -t _YOUR_DOCKER_USERNAME_/robovsdino .`
3. After all downloads and operations finished, run the command: `docker run -p 8080:8080 _YOUR_DOCKER_USERNAME_/robovsdino`

## How to use
For both local machine run and docker run, after you completed steps, you can run the simulation: `./simulate.sh` .
Simulation guides you for all features but if you would like to send personal requests, here is the explanation.

### Board
`http(s)://_YOUR_URL_:PORT/api/v1/board/create` Creates new 50x50 board. This app only accepts square boards and you
can change board by changing variable in **consdef.go** file. If you need to change the size of the board, you must 
change the Column map too.

`http(s)://_YOUR_URL_:PORT/api/v1/board/status` Shows detailed board status.

### Robot
`http(s)://_YOUR_URL_:PORT/api/v1/robot/create/:row/:column/:face` Creates a robot on specific cell where **row** is
integer variable between LOWERBOUND and UPPERBOUND, **column** is English alphabet letters begin from A and ends with AX 
like MS Excel, **face** is the direction of the robot.

`http(s)://_YOUR_URL_:PORT/api/v1/robot/move/:row/:column/:face` Moves the robot defined in face direction for 1 cell 
where **row** is integer variable between LOWERBOUND and UPPERBOUND, **column** is English alphabet letters begin 
from A and ends with AX  like MS Excel, **face** is the direction of the robot.

`http(s)://_YOUR_URL_:PORT/api/v1/robot/attack/:row/:column/:face` Robot attacks on a dinosaur defined in face direction 
for 1 cell where **row** is integer variable between LOWERBOUND and UPPERBOUND, **column** is English alphabet letters 
begin from A and ends with AX  like MS Excel, **face** is the direction of the robot.

### Dinosaur
`http(s)://_YOUR_URL_:PORT/api/v1/dinosaur/create/:row/:column` Creates a dinosaur on specific cell where **row** is
integer variable between LOWERBOUND and UPPERBOUND, **column** is English alphabet letters begin from A and ends with AX 
like MS Excel

## Known Issues
1. Not all test cases covered.
2. Supports only square boards and if you would like to change board size, you must change Column map too inside 
**constdef.go** file. 
3. The code can be refactored by using POST rather than GET everywhere. For example, robot and dinosaur creation nearly
same except robot has facing direction but it can be handled in one function. I intentionally didn't do it because I
wanted the game to be played by using just a basic web browser rather than `curl` program or specific apps like [Postman]("https://www.postman.com/) .
4. There is no CI/CD pipeline in the solution as deployment place is unknown. As a result, tests are there but may not 
be so meaningful alone. 
5. There are no hard checks about parameters that player supplied. This app assumes that player knows how to play. App won't 
stop because of weird inputs but your requests may return response codes 5xx.
6. There is no long term storage for tracking/auditing. I intentionally left it undeveloped because I didn't want to 
create complications to create folder on file system or mount storage for Docker. As a result, you would not see logging 
mechanism and development in middleware.

## Removing Ambiguity
In the question, it says:
```
a robot can turn left, turn right, move forward, move backward
```
and
```
A robot attack destroys dinosaurs around it (in front, to the left, to the right or behind)
```
Well, front, left, right and behind may change based on where to look to the board so **I replaced as top, left, right, bottom**. 
On the other hand, in the coding perspective, I also enabled directions with cardinal directions/points. All JSON 
responses return `msg` to be shown to user and more fields which are available for developer, not user.

## Code Flow
![EM Challenge Code Flow](https://canartuc.com/grover_robovsdino.png "EM Challenge Code Flow")

If you cannot see it properly, [please click here](https://whimsical.com/5cVkzRSZrc7hqkyWZV3kXA)

ENJOY!


