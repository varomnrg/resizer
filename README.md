# Resizer

Resize your images with CLI command

## Introduction

Resizer is a command-line tool designed to resize images quickly and easily. Whether you need to resize images for web projects, social media, or any other purpose, Resizer provides a convenient way to accomplish this task without the need for complex image editing software.

## Installation

To install Resizer, you can use the `go install` command:
```bash
go install github.com/yourusername/resizer
```
Make sure you have Go installed on your system before running the installation command.

## Usage

Resizer provides a simple command-line interface for resizing images. You can use the following command to resize an image:

```bash
resizer <imagepath> <width> <height>
```

To check the resolution of an image, use the following command:
```bash
resizer check <imagepath>
```
Replace `<imagepath>`, `<width>`, and `<height>` with the appropriate values for your image resizing needs.

## Options

Resizer supports the following options:

-   `check`: Check the resolution of the specified image without resizing it.

## Examples

Here are some examples of how to use Resizer:

To resize an image named `example.jpg` to a width of 800 pixels and a height of 600 pixels:

```bash
resizer example.jpg 800 600
```

To check the resolution of the same image without resizing it:
```bash
resizer check example.jpg
```
