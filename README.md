# filemanager

File management web application.

The name filemanager is a placeholder name, it should probably be changed to something cooler.

Currently developed and tested on Linux, specifically Ubuntu 22.04.

### Features

- Web GUI with file management functionality.
- Browse directories in gallery or table view.
- View file details with support for various file formats.
- Looks and feels like a desktop application.
- Multiple users and multiple main directories.
- Simple configuration via yaml config file.
- Security is built with established and proven libraries.
- Really fast and lean.

### Tech details

- Deployed as a single stand-alone self-contained compiled binary executable file.
- Listens on a single port using HTTP.
- Serves web app and API on the same port, under `/` and `/api` respectively.
- Backend is built with Golang, and Chi.
- Frontend is built with JavaScript, Svelte, and SvelteKit.



## Installation

For now, download and compile source code (see Development, it's very straight forward).



## Development

Languages and tools are installed via [mise][mise].

Install mise.

```bash
curl https://mise.run | sh
```

Install requirements with mise.

```bash
mise install
```

Install go packages.

```bash
cd api
go get
```

Install npm packages.

```bash
cd web
npm install
```

Run dev servers.

```bash
cd api
just
```

```bash
cd web
just
```

Build production executable.

```bash
just build
```

Run production executable.  
It listens on 127.0.0.1:4000.  
The example config file has one user with username demo and password demo, and one file area for the path `/tmp`.

```bash
./main -config config.yml
```



[mise]: https://github.com/jdx/mise
