launch ec2
install softwares

```bash
sudo apt update
sudo apt install git nginx -y
sudo apt install golang-go -y
sudo apt install build-essential gcc libc6-dev

cd 
git clone https://github.com/abhinayjangde/goauth.git
cd goauth
```

we are using neon postgres db

create .env file


## build

```bash

go mod tidy # installing dependencies
go build -o goauth ./cmd

# now we will have ./goauth binary
# test 
./goauth

# it will run you api
```

## systemd service

```bash
sudo nano /etc/systemd/system/goauth.service

# past this conent to 'goauth.service` file

`
[Unit]
Description=GoAuth API
After=network.target

[Service]
User=ubuntu
WorkingDirectory=/home/ubuntu/goauth
ExecStart=/home/ubuntu/goauth/goauth

Restart=always

EnvironmentFile=/home/ubuntu/goauth/.env

[Install]
WantedBy=multi-user.target
`
Ctrl+O
Ctrl+X

sudo systemctl daemon-reload
sudo systemctl enable goauth
sudo systemctl start goauth
sudo systemctl status goauth
```

## reverse proxy

```bash

sudo nano /etc/nginx/sites-available/goauth

# paste this content

`
server {
    listen 80;
    server_name _;
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
`
# create symlink

sudo ln -s \
/etc/nginx/sites-available/goauth \
/etc/nginx/sites-enabled/

# remove default
sudo rm /etc/nginx/sites-enabled/default 
# restart
sudo nginx -t
sudo systemctl restart nginx
```

## ec2 security group

make usre these port are open for incomming requests (inbound rule)
22, 80, 443

## test

You can access your api at http://EC2_IP:/health

## https

```bash

sudo certbot --nginx

```

## from ci/cd

      - name: Create .env
        run: |
          cat > .env << EOF
          PORT=${{ secrets.PORT }}
          DATABASE_URL=${{ secrets.DATABASE_URL }}
          JWT_SECRET=${{ secrets.JWT_SECRET }}
          JWT_EXPIRE_HOURS=${{ secrets.JWT_EXPIRE_HOURS }}
          EOF
