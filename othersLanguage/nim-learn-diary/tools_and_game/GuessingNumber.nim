import random, strutils
echo "Guess number in 1..100!"
let number = rand 1..100
while true:
    echo "Please input your guess:"
    let guess = readLine(stdin).parseInt
    if guess < number: echo "small!" elif guess > number: echo "big!" else: echo "you win!"; break