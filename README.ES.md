# cc-identity-go

## Descripción
`cc-identity-go` es un chaincode liviano implementado en el lenguaje go para la tecnología de blockchain empresarial Hyperledger Fabric. 
Este chaincode gestiona las credenciales de los participantes (emisores, usuarios, etc.) en una red blockchain y controla la ejecución de transacciones en la blockchain mediante la definición de roles y permisos.
Esas credenciales, que pueden contener datos identificativos se almacenan en un wallet de manera descentralizada (controlada por el usuario) o semi-descentralizada (controlada por las organizaciones).

Los certificados de identidad deben ser emitido por una Autoridad de Certificación (AC) previamente registrada en cc-identity-go. Cada organización puede verificar/consultar los atributos de un participante. Un solo DID (identificador) debería ser suficiente para consultar toda la información relacionada con una identidad y sus atributos. Un certificado solo puede ser gestionado por la organización que lo emitió.

Funcionalidades:
- Gestionar nuevas identidades
- Gestionar emisores
- Gestionar roles y permisos
- Validar certificados digitales


## Usar cc-identity-go como un módulo (de tipo contrato inteligente) dentro de un chaincode

```bash
go get github.com/kmilodenisglez/cc-identity-go
```

## Usar cc-identity-go como un servicio (chaincode)

### Empaquetar e instalar

Empaquete e instale el chaincode en nodo-pares con los siguientes comandos:

```
tar cfz code.tar.gz connection.json
tar cfz external-chaincode.tgz metadata.json code.tar.gz

peer lifecycle chaincode install external-chaincode.tgz
```

Ejecute el siguiente comando para consultar el ID del paquete del chaincode que acaba de instalar:

```
peer lifecycle chaincode queryinstalled -o $ORDERER_ADDRESS --tlsRootCertFiles $ORDERER_TLS_CA
```

El comando devolverá un resultado similar al siguiente:
```bash
Chaincode code package identifier: ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5

# Installed chaincodes on peer:
Package ID: ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5
Label: ccidentity_1.0
```

Copie el ID del paquete de chaincode devuelto en una variable de entorno para usar en comandos posteriores (su ID puede ser diferente):

```bash
# en linux o darwin use export, ej:
export CHAINCODE_ID=ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5

# en Windows debe usar set, ej:
set CHAINCODE_ID=ccidentity_1.0:5beffad255c6af3744d419ae7b978aa9609386f6c81dd98184808746cea399d5
```

### Contruir

Para construir el chaincode:


```
# linux o darwin
go build -o ccass_identity_binary

# windows
go build -o ccass_identity_binary.exe
```

Establezca la dirección donde va a estar ejecutándose el chaincode:

```
# linux
export CHAINCODE_SERVER_ADDRESS=127.0.0.1:9999

# windows
set CHAINCODE_SERVER_ADDRESS=127.0.0.1:9999
```
### Iniciar

Inicie el servicio de chaincode:

```
# linux
./ccass_identity_binary

# windows
ccass_identity_binary.exe
```
## Activar

aprobar y confirmar el chaincode:

```
peer lifecycle chaincode approveformyorg -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA --channelID $CHANNEL_NAME --name $CC_NAME --version 1 --package-id $CHAINCODE_ID --sequence 1

peer lifecycle chaincode commit -o $ORDERER_ADDRESS --tls --cafile $ORDERER_TLS_CA --channelID $CHANNEL_NAME --name $CC_NAME --version 1 --sequence 1
```

## Uso del administrador de procesos PM2

Recomendamos usar PM2 para administrar el proceso del chaincode, usar solo cuando emplee el chaincode como un servicio.

[que es pm2?](https://pm2.keymetrics.io/docs/usage/process-management/)

### Iniciar cc-identity-go usando PM2

```bash
pm2 start ccass_identity_binary
```
<br>

### Detener cc-identity-go usando PM2

```bash
$ pm2 stop ccass_identity_binary