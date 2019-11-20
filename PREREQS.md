
## Hyperledger Fabric 개발 환경 구축

- 필수 구성 요소 다운로드 및 설치  
  
```
curl -O https://hyperledger.github.io/composer/latest/prereqs-ubuntu.sh
chmod u+x prereqs-ubuntu.sh
./prereqs-ubuntu.sh
```
  
  
- sudo 없이 docker 명령어 실행  

```
sudo usermod -aG docker $USER
```
  
  
 - 서버 재실행  
    
```
sudo reboot
```
  

 - java 설치 (JDK, JRE)  

```
sudo apt-get install openjdk-8-jre
sudo apt-get install openjdk-8-jdk
```

- maven 설치

```
sudo apt install maven
```

