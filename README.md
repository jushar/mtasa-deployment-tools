# MTA:SA Server Deployment Tools
This repository contains a bunch of tools and code snippets that are useful for deploying MTA:SA server scripts in a production environment.
The components don't work out of the box, but are just examples.

## License
Everything available in this repository is license under MIT (see _LICENSE_ in the top directory).

## Components
### Workerunit
The _Workerunit_ is a tool writen in _Golang_ to control and manage the MTA:Server via a RESTful API while it is running.
This tool is preferably used in a _Docker_ environment.

Supported features are:
* Start/Stop/Restart
* Read logs (adjustable ring buffer)
* Execute commands
* Watchdog (automatically restart on crashes)
* Patch config entries based upon environment variables

### Dockerenv
Contains an example _Dockerfile_ that packs an MTA:SA server and adds scripts and the _Workerunit_ to it.

### Buildenv
An example _.gitlab-ci.yml_ (GitLab CI configuration).
