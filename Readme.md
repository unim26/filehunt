# filehunt

`filehunt` is a lightning-fast, cross-platform command-line utility written in Go designed to locate files across your directory structures instantly.

## Features
* Speed: Traverses directories efficiently using modern standard libraries.
* Simplicity: Clean, intuitive command-line interface for everyday use.
* Standalone Binary: Runs anywhere on your system with zero dependencies.

## Installation

Option 1: Download the binary (Recommended)
Visit the Releases (https://github.com/unim26/filehunt/releases) page and download the executable for your operating system (Windows, macOS, or Linux). Extract it and add it to your system PATH.

Option 2: Install via Go
If you have Go installed on your machine, you can build and install it directly to your system bin folder:

```
go install github.com/unim26/filehunt@latest
```

## Usage

Run filehunt by providing the target directory and the file you are looking for.

Basic Search:
filehunt [directory_path] [filename]

(For complete usages instruction check out the usages.md file or use command ```filehunt -h``` or ```filehunt --help```)

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this tool better, please fork the repository and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Roadmap
* Implement concurrent searching for faster system-wide results.
* Add command-line flags for case-insensitive matching.
* Support for regular expression searches.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

---
Built by Abhishek Kumar.
