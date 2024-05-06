ASCII Art Banner Generator is a simple command-line tool written in Go for generating ASCII art banners.

# Usage
Run the program: 
student$ go run . "Your Text Here" [optional: path/to/banner/file.txt] | cat -e

Example:
--- go run . "Hello World" thinkertoy.txt | cat -e

If no path to the banner file is provided, the program will use the default banner file ("standard.txt").

# Features

- Generates ASCII art banners from text input.
- Supports custom banner files.
- Checks if all characters in the input text are within the printable ASCII range.

# Contact

If you have any questions or suggestions, feel free to contact me at [lahmami0020@gmail.com].
