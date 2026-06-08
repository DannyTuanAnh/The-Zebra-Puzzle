# The Zebra Puzzle (Einstein's Riddle)

This is a program to solve **The Zebra Puzzle**, a famous logic puzzle also known as **Einstein's Riddle**.

## Description

This puzzle requires you to answer "Who owns the Zebra? Who drinks Water?" based on a set of logical clues.

**5 Houses** with the following characteristics:

- **Colors**: Yellow, Blue, Red, Green, Ivory
- **Nationalities**: Britt, Spanish, Ukrainian, Norwegian, Japanese
- **Pets**: Horse, Dog, Snails, Fox, Zebra
- **Drinks**: Coffee, Tea, Milk, Orange Juice, Water
- **Tobaccos**: Old Gold, Kools, Chesterfields, Parliaments, Lucky Strike

## How to Run the Program

### Requirements

- Go 1.16 or higher if running directly from Go source files

### Method 1: Run the Executable File

#### On Windows

1. Open Command Prompt or PowerShell
2. Navigate to the directory containing the executable file
3. Run the command:

```bash
.\the-zebra-puzzle.exe
```

Or simply: Double-click the `the-zebra-puzzle.exe` file

#### On Mac

1. Open Terminal
2. Navigate to the directory containing the executable file (replace with your actual path):

```bash
cd ~/path/to/The\ Zebra\ Puzzle
```

Example: If the file is in the Downloads folder:

```bash
cd ~/Downloads/The\ Zebra\ Puzzle
```

3. Grant execute permission (only need to do this once):

```bash
chmod +x the-zebra-puzzle
```

4. Run the program:

```bash
./the-zebra-puzzle
```

### Method 2: Run Directly from Go

```bash
go run .
```

or

```bash
go run main.go
```

## Output

The program will:

1. **Apply AC3** - Reduce the domain of variables by removing incompatible values
2. **Use Backtracking Search with MAC + MRV + Degree + LCV** - Find a solution that satisfies all constraints
3. **Display Results** - Show variables sorted by house position (1-5)

### Example Output:

```
Domain of each variable after using AC3:
[Shows the domain of each variable after reduction]

Using MAC in backtrack
Result:
Kools:1 Yellow:1 Norwegian:1 Water:1 Fox:1
Tea:2 Ukrainian:2 Horse:2 Chesterfields:2 Blue:2
Britt:3 Red:3 Snails:3 OldGold:3 Milk:3
Spaniard:4 OrangeJuice:4 LuckyStrike:4 Dog:4 Ivory:4
Japanese:5 Coffee:5 Zebra:5 Parliaments:5 Green:5

The owner of the zebra lives in house 5.
The owner who drinks water lives in house 1.
```
