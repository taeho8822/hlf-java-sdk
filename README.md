
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
## Hyperledger Fabric 네트워크 구축

- git 다운로드


	```
	$ git clone https://github.com/taeho8822/hlf-java-sdk.git
	```

- 네트워크 설치 

   ```
   cd network
   chmod +x build.sh
   ./build.sh
   ```

- 네트워크 stop

   ```
   cd network
   chmod +x stop.sh
   ./stop.sh
   ```

- 네트워크 삭제

   ```
   cd network
   chmod +x teardown.sh
   ./teardown.sh
   ```

- JAVA SDK 메이븐 빌드 & 패킷

   ```
   cd ../java
   mvn install
   ```

- blockchain-client.jar 으로 복사

   ```
   cd target
   cp blockchain-java-sdk-0.0.1-SNAPSHOT-jar-with-dependencies.jar blockchain-client.jar
   ```

- network_resources 폴더로 복사

   ```
   cp blockchain-client.jar ../../network_resources
   ```


- 채널생성 `mychannel`

   ```
   cd ../../network_resources
   java -cp blockchain-client.jar org.example.network.CreateChannel
   ```

Output:

   ```Apr 20, 2018 5:11:42 PM org.example.util.Util deleteDirectory
      INFO: Deleting - users
      Apr 20, 2018 5:11:45 PM org.example.network.CreateChannel main
      INFO: Channel created mychannel
      Apr 20, 2018 5:11:45 PM org.example.network.CreateChannel main
      INFO: peer0.org1.example.com at grpc://localhost:7051
      Apr 20, 2018 5:11:45 PM org.example.network.CreateChannel main
      INFO: peer1.org1.example.com at grpc://localhost:7056
      Apr 20, 2018 5:11:45 PM org.example.network.CreateChannel main
      INFO: peer0.org2.example.com at grpc://localhost:8051
      Apr 20, 2018 5:11:45 PM org.example.network.CreateChannel main
      INFO: peer1.org2.example.com at grpc://localhost:8056
   ```


- fabcar 체인코드 배포

   ```
   java -cp blockchain-client.jar org.example.network.DeployInstantiateChaincode
   ```

   Output:

   ```Apr 23, 2018 10:25:22 AM org.example.client.FabricClient deployChainCode
      INFO: Deploying chaincode fabcar using Fabric client Org1MSP admin
      Apr 23, 2018 10:25:22 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code deployment SUCCESS
      Apr 23, 2018 10:25:22 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code deployment SUCCESS
      Apr 23, 2018 10:25:22 AM org.example.client.FabricClient deployChainCode
      INFO: Deploying chaincode fabcar using Fabric client Org2MSP admin
      Apr 23, 2018 10:25:22 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code deployment SUCCESS
      Apr 23, 2018 10:25:22 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code deployment SUCCESS
      Apr 23, 2018 10:25:22 AM org.example.client.ChannelClient instantiateChainCode
      INFO: Instantiate proposal request fabcar on channel mychannel with Fabric client Org2MSP admin
      Apr 23, 2018 10:25:22 AM org.example.client.ChannelClient instantiateChainCode
      INFO: Instantiating Chaincode ID fabcar on channel mychannel
      Apr 23, 2018 10:25:25 AM org.example.client.ChannelClient instantiateChainCode
      INFO: Chaincode fabcar on channel mychannel instantiation java.util.concurrent.CompletableFuture@723ca036[Not completed]
      Apr 23, 2018 10:25:25 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code instantiation SUCCESS
      Apr 23, 2018 10:25:25 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code instantiation SUCCESS
      Apr 23, 2018 10:25:25 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code instantiation SUCCESS
      Apr 23, 2018 10:25:25 AM org.example.network.DeployInstantiateChaincode main
      INFO: fabcar- Chain code instantiation SUCCESS
   ```

   > **Note:** The chaincode fabcar.go was taken from the fabric samples available at - https://github.com/hyperledger/fabric-samples/tree/release-1.4/chaincode/fabcar/go.

- enroll user 등록

   ```
   java -cp blockchain-client.jar org.example.user.RegisterEnrollUser
   ```

   Output:

   ```Apr 23, 2018 10:26:34 AM org.example.util.Util deleteDirectory
      INFO: Deleting - users
      log4j:WARN No appenders could be found for logger (org.hyperledger.fabric.sdk.helper.Config).
      log4j:WARN Please initialize the log4j system properly.
      log4j:WARN See https://logging.apache.org/log4j/1.2/faq.html#noconfig for more info.
      Apr 23, 2018 10:26:35 AM org.example.client.CAClient enrollAdminUser
      INFO: CA -http://localhost:7054 Enrolled Admin.
      Apr 23, 2018 10:26:35 AM org.example.client.CAClient registerUser
      INFO: CA -http://localhost:7054 Registered User - user1524459395783
      Apr 23, 2018 10:26:36 AM org.example.client.CAClient enrollUser
      INFO: CA -http://localhost:7054 Enrolled User - user1524459395783
   ```

 - Perform Invoke and Query on network

   ```
   java -cp blockchain-client.jar org.example.chaincode.invocation.InvokeChaincode
   ```

   Output:

   ```Apr 20, 2018 5:13:03 PM org.example.client.CAClient enrollAdminUser
     INFO: CA -http://localhost:7054 Enrolled Admin.
     Apr 20, 2018 5:13:04 PM org.example.client.ChannelClient sendTransactionProposal
     INFO: Sending transaction proposal on channel mychannel
     Apr 20, 2018 5:13:04 PM org.example.client.ChannelClient sendTransactionProposal
     INFO: Transaction proposal on channel mychannel OK SUCCESS with transaction
     id:a298b9e27bdb0b6ca18b19f9c78a5371fb4d9b8dd199927baf37379537ca0d0f
     Apr 20, 2018 5:13:04 PM org.example.client.ChannelClient sendTransactionProposal
     INFO:
     Apr 20, 2018 5:13:04 PM org.example.client.ChannelClient sendTransactionProposal
     INFO: java.util.concurrent.CompletableFuture@22f31dec[Not completed]
     Apr 20, 2018 5:13:04 PM org.example.chaincode.invocation.InvokeChaincode main
     INFO: Invoked createCar on fabcar. Status - SUCCESS
  ```

   ```
   java -cp blockchain-client.jar org.example.chaincode.invocation.QueryChaincode
   ```

   Output:

   <pre>
    Apr 20, 2018 5:13:28 PM org.example.client.CAClient enrollAdminUser
    INFO: CA -http://localhost:7054 Enrolled Admin.
    Apr 20, 2018 5:13:29 PM org.example.chaincode.invocation.QueryChaincode main
    INFO: <b>Querying for all cars ...</b>
    Apr 20, 2018 5:13:29 PM org.example.client.ChannelClient queryByChainCode
    INFO: Querying queryAllCars on channel mychannel
    Apr 20, 2018 5:13:29 PM org.example.chaincode.invocation.QueryChaincode main
    INFO: <b>[{"Key":"CAR1", "Record":{"make":"Chevy","model":"Volt","colour":"Red","owner":"Nick"}}]</b>
    Apr 20, 2018 5:13:39 PM org.example.chaincode.invocation.QueryChaincode main
    INFO: <b>Querying for a car - CAR1</b>
    Apr 20, 2018 5:13:39 PM org.example.client.ChannelClient queryByChainCode
    INFO: Querying queryCar on channel mychannel
    Apr 20, 2018 5:13:39 PM org.example.chaincode.invocation.QueryChaincode main
    INFO: <b>{"make":"Chevy","model":"Volt","colour":"Red","owner":"Nick"}</b>
   </pre>



## License
This code pattern is licensed under the Apache Software License, Version 2.  Separate third party code objects invoked within this code pattern are licensed by their respective providers pursuant to their own separate licenses. Contributions are subject to the [Developer Certificate of Origin, Version 1.1 (DCO)](https://developercertificate.org/) and the [Apache Software License, Version 2](https://www.apache.org/licenses/LICENSE-2.0.txt).

[Apache Software License (ASL) FAQ](https://www.apache.org/foundation/license-faq.html#WhatDoesItMEAN)

