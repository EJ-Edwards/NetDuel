import os
from py_compile import main
import random
from player import Player

def clear_screen():
    os.system('cls' if os.name == 'nt' else 'clear')

def determine_winner(p1, p2):
    if p1.move == p2.move:
        print(f"Both players selected {p1.move}. It's a tie!")
        return None
    elif (
        (p1.move == "rock" and p2.move == "scissors") or
        (p1.move == "paper" and p2.move == "rock") or
        (p1.move == "scissors" and p2.move == "paper")
    ):
        print(f"{p1.name} wins! {p1.move.capitalize()} beats {p2.move}!")
        return p1
    else:
        print(f"{p2.name} wins! {p2.move.capitalize()} beats {p1.move}!")
        return p2

def pvp_game():
    player1 = Player(input("Enter Player 1 name: "))
    player2 = Player(input("Enter Player 2 name: "))
    lives = 3
    player1.score = lives
    player2.score = lives
    round_num = 1
    while player1.score > 0 and player2.score > 0:
        print(f"\n--- Round {round_num} ---")
        player1.choose_move()
        clear_screen()
        player2.choose_move()
        print(f"{player1.name} chose {player1.move}, {player2.name} chose {player2.move}.")
        winner = determine_winner(player1, player2)
        if winner == player1:
            player2.score -= 1
        elif winner == player2:
            player1.score -= 1
        print(f"Lives: {player1.name} = {player1.score}, {player2.name} = {player2.score}")
        round_num += 1
    if player1.score > player2.score:
        print(f"\n{player1.name} wins the game!")
    else:
        print(f"\n{player2.name} wins the game!")

def single_player_game():
    player = Player(input("Enter your name: "))
    computer = Player("Computer")
    lives = 3
    player.score = lives
    computer.score = lives
    round_num = 1
    last_player_move = None
    while player.score > 0 and computer.score > 0:
        print(f"\n--- Round {round_num} ---")
        player.choose_move()
        computer.ai_choose_move(last_player_move)
        print(f"{player.name} chose {player.move}, Computer chose {computer.move}.")
        winner = determine_winner(player, computer)
        if winner == player:
            computer.score -= 1
        elif winner == computer:
            player.score -= 1
        print(f"Lives: {player.name} = {player.score}, Computer = {computer.score}")
        last_player_move = player.move
        round_num += 1
    if player.score > computer.score:
        print(f"\n{player.name} wins the game!")
    else:
        print(f"\nComputer wins the game!")



if __name__ == "__main__":
    main()