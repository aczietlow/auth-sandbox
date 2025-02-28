# Auth Sandbox

Make a sandbox to better explore and understand OAuth

## How to use

To run locally with live reloading

> `air`

Then visit `http://localhost:1322`

## Architecture

- /cmd - main application for this project
  - Here be the web server (echo in this case)
- /pkg - library code that's okay for use in external applications. 
- /tmp - build assets used during compiling and serving the website
- /website - the project's website used as a PoC implementation

## What this is NOT

HTMX, Echo, and Air were all new to me during this project. While I went as far as reading the docs, there were a means to an end.

- [HTMX](https://htmx.org/) - dynamic web using [HATEOAS](https://en.wikipedia.org/wiki/HATEOAS)
-https://pkg.go.dev/html/template [Echo](https://echo.labstack.com/docs) - popular http web server
- [Air](https://github.com/air-verse/air) - live reloading utility to improve local dev 
- [go's html/template](https://pkg.go.dev/html/template) - generate html templates.
  - This is somewhat familiar to me from Hugo.

