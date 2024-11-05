# Checkers-Torram Blockchain

A decentralized checkers game built with Cosmos SDK v0.50. This application demonstrates a complete blockchain implementation with custom message handling and game state management.

## Prerequisites

Before you begin, ensure you have:
- Go 1.21+
- Make
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/pratikasr/checkers-torram.git
cd checkers-torram
```

2. Build and install:
```bash
make install
```

3. Initialize and start the chain:
```bash
./scripts/setup_chain.sh
```

## Usage

### Creating a Game

```bash
# Export account addresses
export ALICE=$(checkd keys show alice -a --keyring-backend test)
export BOB=$(checkd keys show bob -a --keyring-backend test)

# Create a game
checkd tx checkerstorram create-game $ALICE $BOB \
--from alice \
--keyring-backend test \
--chain-id checkers-1 \
--fees 6000stake
```

### Querying a Game

```bash
# Query using game index (returned from create-game transaction)
checkd query checkerstorram get-game 0 --output json

# Example Response:
{
  "storedGame": {
    "index": "0",
    "board": "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
    "turn": "b",
    "black": "cosmos1...", // ALICE address
    "red": "cosmos1...",   // BOB address
    "game_start_time": "1699167812",
    "game_end_time": "0"
  }
}
```

### Export Chain State

To export the entire chain state including all games:

```bash
# Export chain state to a JSON file
checkd export --output-document checkd_export.json

# Open the exported file
open checkd_export.json
```

The export will contain all game states, including timestamps and player information.

### Game State Structure

Each game stores:
```typescript
{
    "index": string,
    "board": string,
    "turn": string,
    "black": string,
    "red": string,
    "game_start_time": int64,
    "game_end_time": int64
}
```

## Technical Specification

### Custom Service Names
- Message Service: `CheckersTorram`
- RPC Method: `CheckersCreateGm`
- Request Type: `ReqCheckersTorram`
- Response Type: `ResCheckersTorram`

### Directory Structure
```
.
├── app/            # Application initialization
├── cmd/            # Command line interface
├── proto/          # Protobuf definitions
├── x/             # Modules
│   └── checkerstorram/  # Main game module
└── scripts/        # Helper scripts
```

## Development

### Building
```bash
make build
```

### Testing
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Commands Available
```bash
make install    # Install the binary
make clean      # Clean build files
make test      # Run tests
```