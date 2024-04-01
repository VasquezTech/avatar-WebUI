
### Install docker
#### From [docs.docker](https://docs.docker.com/engine/install/) for Ubuntu
##### Remove old:
```
 for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
```
##### Setup apt:
```
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

##### Install:
```
 sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```
##### Manage Docker as a non-root user:
```
sudo groupadd docker
sudo usermod -aG docker $USER
```
##### start on boot:
```
sudo systemctl enable docker.service
sudo systemctl enable containerd.service

```
### Avatar-ui startup
```
make
```
 
Open `https:/localhost/8065`

#### Avatar traits collected from: https://www.avatarsinpixels.com/
