import random
while true:
    echo "请输入 rock/paper/scissors"
    let random = rand 1..3
    let computer = if random == 1: "rock" elif random == 2: "paper" else: "scissors"
    case readLine stdin
    of "rock":
        case computer
        of "rock":     echo "计算机为", computer, "，平局\n"
        of "paper":    echo "计算机为", computer, "，获胜\n"
        of "scissors": echo "计算机为", computer, "，败北\n"
    of "paper":
        case computer
        of "rock":     echo "计算机为", computer, "，败北\n"
        of "paper":    echo "计算机为", computer, "，平局\n"
        of "scissors": echo "计算机为", computer, "，获胜\n"
    of "scissors":
        case computer
        of "rock":     echo "计算机为", computer, "，获胜\n"
        of "paper":    echo "计算机为", computer, "，败北\n"
        of "scissors": echo "计算机为", computer, "，平局\n"