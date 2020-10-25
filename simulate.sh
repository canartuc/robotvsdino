#!/usr/bin/env bash

BOARD_SIZE=49
API_ROOT_URL="http://127.0.0.1:8080/api/v1"
CURL=$(which curl)

# Create board
printf "\n(***) Create $((${BOARD_SIZE} + 1)) x $((${BOARD_SIZE} + 1)) board\n"
${CURL} -XGET ${API_ROOT_URL}/board/create
printf "\n"
sleep 15

# Create robot at row 5 column C with left face direction
printf "\n(***) Create robot at row 5 column C with left face direction\n"
${CURL} -XGET ${API_ROOT_URL}/robot/create/5/c/left
printf "\n"
sleep 15

# Try to re-create board again - It should not create
printf "\n(***) Re-create the board again - It should not create and give warning\n"
${CURL} -XGET ${API_ROOT_URL}/board/create
printf "\n"
sleep 15

# Re-create robot at row 5 column C with left face direction - It should not create
printf "\n(***) Re-create robot at row 5 column C with left face direction - It should not create and give warning\n"
${CURL} -XGET ${API_ROOT_URL}/robot/create/5/c/left
printf "\n"
sleep 15

# Create Dinosaur at row 5 column C - It should not create
printf "\n(***) Create dinosaur at row 5 column C - It should not create and give warning\n"
${CURL} -XGET ${API_ROOT_URL}/dinosaur/create/5/c
printf "\n"
sleep 15

# Move robot to one cell right
printf "\n(***) Move robot at row 5 column C with left face direction to 1 cell right\n"
${CURL} -XGET ${API_ROOT_URL}/robot/move/5/d/right
printf "\n"
sleep 15

# Try re-create dinosaur at row 5 column C - It should create
printf "\n(***) Create dinosaur at row 5 column C - It should create because robot moved\n"
${CURL} -XGET ${API_ROOT_URL}/dinosaur/create/5/c
printf "\n"
sleep 15

# Move robot to one cell left, there is a dinosaur - Robot should not move
printf "\n(***) Move robot to 1 cell left as row 5 column C - Robot should not move because there is dinosaur\n"
${CURL} -XGET ${API_ROOT_URL}/robot/move/5/c/left
printf "\n"
sleep 15

# Attack robot to one cell left, there is a dinosaur - Robot should attack
#printf "\n(***) Attack robot to 1 cell left as row 5 column C facing left - Robot should attack and no more dinosaur\n"
#${CURL} -XGET ${API_ROOT_URL}/robot/attack/5/c/left
#printf "\n"
#sleep 15

# Status should have 2499 empty cells and 1 robot at row 5, column C, face left and 0 dinosaur
printf "\n(***) Status should have 2499 empty cells and 1 robot at row 5, column C, face left and 0 dinosaur\n"
${CURL} -XGET ${API_ROOT_URL}/board/status
printf "\n"
sleep 15