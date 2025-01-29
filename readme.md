# Go Slot Machine Game 🎰

A simple command-line slot machine game implemented in Go. Players can bet virtual money and try their luck at winning by matching symbols across different lines.

## Features

- Interactive command-line interface
- Multiple symbol types with different winning multipliers
- Balance tracking system
- Three-row by three-column slot machine grid
- Line-based winning combination detection
- Customizable betting amounts

## Game Symbols and Multipliers

The game features four different symbols with varying frequencies and payouts:

- Symbol 'A': 20x multiplier (4 instances in reel)
- Symbol 'B': 10x multiplier (7 instances in reel)
- Symbol 'C': 5x multiplier (12 instances in reel)
- Symbol 'D': 2x multiplier (20 instances in reel)

## Prerequisites

Before running the game, ensure you have:

- Go installed on your machine (version 1.16 or higher recommended)
- Git for cloning the repository

## Installation

1. Clone the repository:

```bash
git clone [repository-url]
cd go-slot-machine
```

2. Build the game:

```bash
go mod init casino

```

3. Run the game:

```bash
go run .
```

## How to Play

1. When you start the game, you'll be prompted to enter your name
2. You start with a balance of $200
3. Each round:
   - Enter your bet amount (or 0 to quit)
   - The slot machine will spin and display the results
   - Winning combinations are automatically calculated
   - Your balance is updated based on wins/losses
4. The game continues until you either:
   - Run out of money
   - Choose to quit by entering 0 as your bet

## Project Structure

```
go-slot-machine/
├── main.go          # Main game loop and initialization
├── welcome.go       # Welcome screen and bet handling
├── spin.go         # Slot machine spinning mechanics
└── win.go          # Win condition checking logic
```

## Game Rules

- A winning line consists of three matching symbols in a row
- Each row is checked independently for wins
- Winnings are calculated by multiplying your bet by the symbol's multiplier
- You can't bet more than your current balance
- The game ends when you run out of money or choose to quit

## Development

To modify the game, you can adjust the following in `main.go`:

- Initial balance (default: 200)
- Symbol frequencies
- Multiplier values
- Grid dimensions (currently 3x3)

## Contributing

Feel free to fork the repository and submit pull requests with improvements or bug fixes.

## License

[Add your chosen license here]
