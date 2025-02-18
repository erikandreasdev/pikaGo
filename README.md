I'll create a comprehensive README.md file for your pikaGo project:

```markdown
# pikaGo

A lightweight system tray application for controlling screen brightness in Linux systems using X11.

## Features

- Simple system tray interface
- Quick brightness adjustments (25%, 50%, 75%, 100%)
- Works with any X11-compatible display
- Minimal resource usage

## Prerequisites

- Linux system with X11
- Go 1.19 or higher
- `xrandr` utility installed
- GTK3 and its development files

For Ubuntu/Debian-based systems, install dependencies with:
```bash
sudo apt-get install libgtk-3-dev libappindicator3-dev
```

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/yourusername/pikaGo.git
cd pikaGo
```

2. Build the application:
```bash
go build
```

3. Run the application:
```bash
./pikaGo
```

### Installing from .deb package

Download the latest .deb package from the releases page and install it:

```bash
sudo dpkg -i pikaGo_1.0-1.deb
```

## Building the .deb package

To create a .deb package:

```bash
# Create build directories
mkdir -p build/deb/pikaGo_1.0-1/DEBIAN
mkdir -p build/deb/pikaGo_1.0-1/usr/local/bin

# Copy files and create package
cp pikaGo build/deb/pikaGo_1.0-1/usr/local/bin/

# Build the application
go build -o pikaGo

# Copy binary
cp pikaGo build/deb/pikaGo_1.0-1/usr/local/bin/

# Create control file
cat > pikaGo_1.0-1/DEBIAN/control << EOF
Package: pikaGo
Version: 1.0-1
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Your Name <your.email@example.com>
Description: Brightness Control Application
 A system tray application for controlling screen brightness
Depends: libgtk-3-0, libappindicator3-1
EOF

# Set permissions
chmod -R 755 pikaGo_1.0-1

# Build the package
dpkg-deb --build build/deb/pikaGo_1.0-1
```

## Usage

After launching pikaGo, you'll see a system tray icon. Click on it to access brightness controls:

- 100%: Full brightness
- 75%: Three-quarters brightness
- 50%: Half brightness
- 25%: Quarter brightness
- Quit: Exit the application

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [systray](https://github.com/getlantern/systray) - For the system tray implementation
- The X.Org project for xrandr
