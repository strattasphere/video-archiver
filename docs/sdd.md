# Software Design Document

## System Overview

This is a simple frontend to using yt-dlp scripts to archive videos onto a local or remote machine. It runs in a container and is accessible through a web app.

## Quick Notes 

- A web application
- Runs in a container
- Downloads content into a non-networked storage location (NFS, Samba) in the case that it fails
- Move downloaded files to network storage location when it is connected
- Command-line interface:
  - Run server
