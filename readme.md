# Personal Wallet CLI App

## Overview

Personal Wallet CLI App, written in Go, empowers you to manage your Solana wallet directly from your command line. It utilizes the Cobra CLI framework for a user-friendly experience, offering functionalities like:

* Creating new wallets
* Checking balances (future implementation)
* Requesting testnet SOL airdrops (limited to Devnet)
* Transferring SOL to other wallets (future implementation)

## Features

* **Create New Wallets:** Generate a secure Solana keypair and store the private key safely. (**Crucial Note:** Your private key grants access to your wallet and funds. Employ a password manager or hardware wallet for enhanced security.)
* **Check Wallet Balance (Future Implementation):** Retrieve the balance associated with your wallet's public key.
* **Request Testnet SOL Airdrops (Devnet Only):** Obtain testnet SOL for development and testing purposes. (**Disclaimer:** Airdrops are unavailable on the mainnet.)
* **Transfer SOL to Other Wallets (Future Implementation):** Send SOL from your wallet to another wallet's public key.

## Installation

**Prerequisites:**

* **Go Programming Language:** Ensure you have Go installed on your system (version 1.18 or later recommended). Download and install it from the official website: https://go.dev/doc/install
* **Cobra CLI Tool (Optional):** While not mandatory, the Cobra CLI tool simplifies command creation within Go applications. Install it using the instructions on the GitHub repository: https://github.com/spf13/cobra/blob/main/cobra.go

**Steps:**

1. **Clone the Repository:**

   ```bash
   git clone [https://github.com/denniskamara/personal-wallet.git](https://github.com/denniskamara/personal-wallet.git)
Use code with caution.
Navigate to the Project Directory:

Bash
cd personal-wallet
Use code with caution.
Build and Install the CLI App:

Bash
go install
Use code with caution.
This creates an executable file named personal-wallet (or similar) in your system's default Go binary location.

Usage
The CLI app provides several commands:

createWallet: Creates a new Solana wallet and stores its private key securely.
importWallet (Future Implementation): Import an existing wallet using its private key (planned for a future version).
requestAirdrop: Requests a testnet SOL airdrop for your wallet (limited to Devnet environment).
transfer (Future Implementation): Transfer SOL to another wallet (planned for a future version).
Running Commands:

To use a command, type it followed by any necessary options:

Bash
personal-wallet <command> [options]
Use code with caution.
Example (Creating a New Wallet):

Bash
personal-wallet createWallet
Use code with caution.
Contributing
We welcome contributions! For guidelines, refer to the CONTRIBUTING.md file (if provided).

License
This project is licensed under the MIT License: https://opensource.org/license/mit

Disclaimer
Important: This application is for educational purposes only. Avoid using it on the mainnet with real funds without thorough testing and a deep understanding of cryptocurrency security practices. Always back up your private key securely.

Explanation of cmd package functions (for reference, within the code):

The cmd package contains functions that implement the core functionalities of the CLI app. Here's a brief overview:

CreateNewWallet(RPCEndpoint string) Wallet: Generates a new Solana keypair, stores the private key securely, and creates a Wallet struct for interacting with the network.
ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error): Attempts to import an existing wallet using its private key.
GetBalance(privateKey []byte) (uint64, error): Retrieves the balance associated with a wallet's public key.
RequestAirDrop(privateKey []byte, amount uint64) (string, error): Requests a testnet SOL airdrop for a wallet (limited to Devnet).
Transfer(privateKey []byte, receiver string, amount uint64) (string, error): Transfers SOL from one wallet to another (future implementation).
Explanation of go.mod file (for reference, within the code):

The go.mod file declares the Go module name (personal-wallet), the minimum required Go version (go 1.21.6), and the project's dependencies.