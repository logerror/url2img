# Screenshot with QR Code Generator -- url to image(png)

This repository provides a Go program that captures a screenshot of a webpage, generates a QR code for the same URL, and overlays the QR code and the URL text on the screenshot. The final image is saved as a PNG file.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Introduction

The goal of this project is to demonstrate how to use Go with `chromedp` to automate web browsing and screenshot capture. Additionally, it uses the `go-qrcode` library to generate a QR code for the webpage URL and overlays both the QR code and the URL text on the captured screenshot.

## Features

- Capture a screenshot of a webpage using headless Chrome.
- Generate a QR code for a specified URL.
- Overlay the QR code and URL text on the screenshot.
- Save the final composite image as a PNG file.

## Requirements

To run this program, you will need:

- [Go](https://golang.org/dl/) 1.16 or later
- [Headless Chrome](https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md)
- Fonts that support Truetype, like DejaVu Sans or any other Truetype font.
- The following Go packages:
    - `github.com/chromedp/chromedp`
    - `github.com/golang/freetype`
    - `github.com/golang/freetype/truetype`
    - `github.com/skip2/go-qrcode`
    - `golang.org/x/image/font`

## Installation

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

2. **Install Headless Chrome**: If you haven't installed Chrome or Chromium yet, you can download it from [here](https://www.google.com/chrome/) or use your package manager.

3. **Get the required Go packages**: Run the following command to install the necessary Go packages:

   ```bash
   go get -u github.com/chromedp/chromedp
   go get -u github.com/golang/freetype
   go get -u github.com/skip2/go-qrcode
   go get -u golang.org/x/image/font
   ```
4. Clone the repository:

```bash
  git clone https://github.com/your-username/screenshot-with-qrcode.git
  cd screenshot-with-qrcode
  ```
5. Build the program:
```bash
  go build -o screenshot_with_qrcode main.go
  ```

## Usage
Run the compiled binary with:

```bash
./screenshot_with_qrcode
```
The program will capture a screenshot of the specified URL, generate a QR code, overlay the QR code and the URL text on the screenshot, and save the final image as screenshot_with_qrcode.png.