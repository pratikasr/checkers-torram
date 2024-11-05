#!/bin/bash
set -e

# Function to log messages
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $1"
}

# Variables
CHAIN_ID="checkers-1"
MONIKER="checkers-node"
BINARY="checkd"  # This should match your binary name
HOME_DIR="$HOME/.checkers-torram"
GENESIS_JSON_PATH="$HOME_DIR/config/genesis.json"
APP_TOML_PATH="$HOME_DIR/config/app.toml"
CONFIG_TOML_PATH="$HOME_DIR/config/config.toml"
VALIDATOR_NAME="alice"  # First player
PLAYER_TWO="bob"       # Second player
TOKEN_DENOM="stake"    # Using default stake denom
VALIDATOR_AMOUNT="1000000000000000000000stake"
PLAYER_AMOUNT="1000000000000000000000stake"
GENTX_AMOUNT="1000000000stake"

log "Starting Checkers blockchain setup..."

# Ensure the binary is in the correct location
if [ ! -f "/usr/local/bin/$BINARY" ]; then
    log "Moving $BINARY to /usr/local/bin..."
    sudo mv ~/go/bin/$BINARY /usr/local/bin/
    if [ $? -ne 0 ]; then
        log "Error: Failed to move $BINARY to /usr/local/bin."
        exit 1
    fi
fi

# Initialize the chain
log "Initializing the chain..."
$BINARY init $MONIKER --chain-id $CHAIN_ID
if [ $? -ne 0 ]; then
    log "Error: Failed to initialize the chain."
    exit 1
fi

# Add keys for both players
log "Adding keys for players..."
$BINARY keys add $VALIDATOR_NAME --keyring-backend test
$BINARY keys add $PLAYER_TWO --keyring-backend test
if [ $? -ne 0 ]; then
    log "Error: Failed to add keys."
    exit 1
fi

# Add genesis accounts
log "Adding genesis accounts..."
$BINARY genesis add-genesis-account $VALIDATOR_NAME $VALIDATOR_AMOUNT --keyring-backend test
$BINARY genesis add-genesis-account $PLAYER_TWO $PLAYER_AMOUNT --keyring-backend test
if [ $? -ne 0 ]; then
    log "Error: Failed to add genesis accounts."
    exit 1
fi

# Create gentx
log "Creating genesis transaction..."
$BINARY genesis gentx $VALIDATOR_NAME $GENTX_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
if [ $? -ne 0 ]; then
    log "Error: Failed to create genesis transaction."
    exit 1
fi

# Update minimum-gas-prices in app.toml
log "Updating minimum gas prices..."
sed -i '' 's/^minimum-gas-prices = ""/minimum-gas-prices = "0.025stake"/' "$APP_TOML_PATH"
if [ $? -ne 0 ]; then
    log "Error: Failed to update minimum gas prices in app.toml."
    exit 1
fi

# Collect genesis transactions
log "Collecting genesis transactions..."
$BINARY genesis collect-gentxs
if [ $? -ne 0 ]; then
    log "Error: Failed to collect genesis transactions."
    exit 1
fi

# Validate genesis file
log "Validating genesis file..."
$BINARY genesis validate-genesis
if [ $? -ne 0 ]; then
    log "Error: Genesis file validation failed."
    exit 1
fi

# Update governance params in genesis.json for faster testing
log "Updating governance parameters in genesis.json..."
sed -i '' 's/"max_deposit_period": ".*"/"max_deposit_period": "60s"/' "$GENESIS_JSON_PATH"
sed -i '' 's/"voting_period": ".*"/"voting_period": "60s"/' "$GENESIS_JSON_PATH"
if [ $? -ne 0 ]; then
    log "Error: Failed to update governance parameters in genesis.json."
    exit 1
fi

# Update app.toml for API and gRPC
log "Updating app.toml..."
sed -i '' 's/enable = false/enable = true/' "$APP_TOML_PATH"
sed -i '' 's/swagger = false/swagger = true/' "$APP_TOML_PATH"
sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' "$APP_TOML_PATH"
if [ $? -ne 0 ]; then
    log "Error: Failed to update app.toml."
    exit 1
fi

# Update config.toml CORS settings
log "Updating CORS settings in config.toml..."
sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/' "$CONFIG_TOML_PATH"
if [ $? -ne 0 ]; then
    log "Error: Failed to update CORS settings in config.toml."
    exit 1
fi

# Export environment variables for easy access
log "Exporting player addresses..."
echo "export ALICE=$($BINARY keys show alice -a --keyring-backend test)" >> ~/.bashrc
echo "export BOB=$($BINARY keys show bob -a --keyring-backend test)" >> ~/.bashrc

# Start the chain
log "Starting the Checkers blockchain..."
$BINARY start

log "Setup complete. If you encounter any issues, please check the logs above."