class Player:
    def __init__(self, name):
        self.name = name
        self.score = 0
        self.move = None

    def choose_move(self):
        while True:
            move = input(f"{self.name}, enter a choice (rock, paper, scissors): ").lower()
            if move in ["rock", "paper", "scissors"]:
                self.move = move
                return
            print("Invalid choice. Please choose rock, paper, or scissors.")
        def ai_choose_move(self, last_player_move=None):
            # Simple AI: tries to counter the player's last move
            import random
            if last_player_move == "rock":
                self.move = "paper"
            elif last_player_move == "paper":
                self.move = "scissors"
            elif last_player_move == "scissors":
                self.move = "rock"
            else:
                self.move = random.choice(["rock", "paper", "scissors"])
            print(f"{self.name} (AI) chose {self.move}.")

