from game import pvp_game, single_player_game, determine_winner, clear_screen
from player import Player
from rules import display_rules

if __name__ == "__main__":
    print("Welcome to the Rock-Paper-Scissors Game!")
    while True:
        print("\nMenu:")
        print("1. Single Player")
        print("2. Multiplayer (PvP)")
        print("3. View Rules")
        print("4. Exit")
        choice = input("Select an option (1-4): ")
        if choice == "1":
            single_player_game()
        elif choice == "2":
            pvp_game()
        elif choice == "3":
            display_rules()
        elif choice == "4":
            print("Goodbye!")
            break
        else:
            print("Invalid choice. Please try again.")