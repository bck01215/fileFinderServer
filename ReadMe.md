# fileFinderServer

## What it does?

Finds all files in provided path (including sub-directories and their files) and displays them in a web gui

![Peek 2022-08-26 09-21](https://user-images.githubusercontent.com/77836661/186913315-8d6c688e-718d-4037-92f4-a7242514814b.gif)

## How to use it?


### Requirements

Before you build anything you need to make sure you have a compatible version of go (go 1.19).
You will also need yarn to build the frontend.

### Use

If you run `make server` this will build the application. Static files will be in `./static` and the executable will be `./fileFinderServer`.
If you run `make run`, it will do everything `make server` does and then start the application.

## Author Information

Authors:

- [Brandon Kauffman](mailto:bck01215@gmail.com)
