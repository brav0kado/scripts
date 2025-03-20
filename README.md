# Gogo - Advanced Enumeration & Scanning Tool 🚀

**Gogo is a modular and flexible enumeration tool written in Go. It is designed for network scanning, automation, and other useful functionalities that will be implemented in future versions. This tool is part of my journey to learn Golang and improve my programming skills while building something practical.** ⚡

## Features ✨

🔍 Scan specific ports or ranges

⏳ Adjustable delay for stealth scanning

🎲 Randomized port scanning to evade detection

⚡ Automation capabilities for efficient recon

🔧 Easily extendable for future features

## Installation 🛠️

Ensure you have Go installed on your system. Then, clone the repository and build the project:

## Usage 🚀

Run the tool using the following syntax:

## Commands

🔎 Scan a Target

### Available Flags 🏴‍☠️

-ports <range>: Specify ports to scan (e.g., 22,80,443 or 1-1000)

-delay <milliseconds>: Add a delay between scans for IDS evasion

-randomize: Randomize port scanning order

### Examples 📌

Scan all ports from 1-1024 on a target:
```bash
gogo scan 192.168.1.1 -ports 1-204
```
Scan only ports 22,80,443 with a 1-second delay:
```bash
gogo scna 192.168.1.1 -ports 22,80,443 -delay 1000
```
Randomized scan on 1-100 with a variable delay:

Roadmap 🏗️

🛰️ Expand enumeration functionalities beyond scanning

🌐 Implement automation features for recon workflows

🔄 Add UDP scanning support

🕵️ Implement proxy support for anonymous scanning

🚀 Include multithreading for faster execution

Disclaimer ⚠️

This tool is being developed as part of my learning process in Golang. 🦍 It is intended for educational purposes only and should be used responsibly and only with permissions. The author is not responsible for any misuse of this software.

Contributing 🤝

Feel free to fork this repository and submit pull requests! Contributions, suggestions, and ideas are always welcome.

License 📜

MIT License