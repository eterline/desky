
# Desky ![Desky Logo](./repo/favicon.svg)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)


       /$$$$$$$                      /$$                
      | $$__  $$                    | $$                
      | $$  \ $$  /$$$$$$   /$$$$$$$| $$   /$$ /$$   /$$
      | $$  | $$ /$$__  $$ /$$_____/| $$  /$$/| $$  | $$
      | $$  | $$| $$$$$$$$|  $$$$$$ | $$$$$$/ | $$  | $$
      | $$  | $$| $$_____/ \____  $$| $$_  $$ | $$  | $$
      | $$$$$$$/|  $$$$$$$ /$$$$$$$/| $$ \  $$|  $$$$$$$
      |_______/  \_______/|_______/ |__/  \__/ \____  $$
                                               /$$  | $$
                                              |  $$$$$$/
                                               \______/ 

- [Frontend](https://github.com/eterline/desky-front) (now is private)
- [Backend](https://github.com/eterline/desky-backend)

My own dashboard for homelab server.
Made for pratice and common usage, no more.


Working with:
- web self-hosted apps
- proxmox
- system
- docker (planned)
- ssh (planned)

## Install

To build project and run:

```bash
make init 

sudo ./app
```


## Settings

Systemd service file.
```
[Unit]
Description=Desky dashboard.
After=network.target


[Service]
Type=simple
User=root

ExecStart=/root/desky/app
WorkingDirectory=/root/desky

Restart=always
RestartSec=30


[Install]
WantedBy=default.target
```

## Screenshoots

## License

[MIT](https://choosealicense.com/licenses/mit/)

