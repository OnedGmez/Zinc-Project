# Zinc-Project

_Interfaz para buscar información de base de datos de correos_

### Pre-requisitos 📋
- [SO linux(funciona wsl)](https://docs.microsoft.com/en-us/windows/wsl/install)
- [Nodejs](https://nodejs.org/en)
- [Golang](https://go.dev/)
- [Git](https://git-scm.com/)
- [openObserve/ZincSearch](https://openobserve.ai/docs/quickstart/#openobserve-cloud)

## Copiar el proyecto 🚀
_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

1. _Hacer fork al repositorio para luego clonarlo_
```
# git clone <URL>
```
2. _Una vez clonado en tu PC, navega hasta el directorio del frontend para ejecutar el comando que instalará todos los paquetes utilizados en el proyecto de **frontend**_
```
# npm install
```
3. _Para hacer pruebas en el frontend, primero debemos levantar el servicio de OpenObserve de forma local con el comando_
```
linux: 
# cd openobserve-v0.7.2-linux-amd64
# ZO_ROOT_USER_EMAIL="example@example.com" ZO_ROOT_USER_PASSWORD="Examplepass" ./openobserve
```
4. _Una vez levantado localmente openobserve, en una nueva terminal debemos navegando hasta el directorio de **backend** para levantar el servidor del API con el comando_
```
linux:
opción 1: # ./main
o
opción 2: # go run main.go
```
5. _Una vez levantados ambos servicios, navegamos hasta el directorio de **indexer** y navegamos al directorio de la versión que queremos correr para ejecutar el comando_
```
linux:
opción 1: # ./indexer
o
opción 2: # go run indexer.go
```
6. _Con el indexer finalizado, podemos utilizar el frontend (aún sin los pasos 3, 4 y 5 se puede utilizar pero no mostrará información), para ello, navegamos al directorio **frontend** y ejecutamos el comando para correrlo localmente_
```
# npm run serve
```

## Construido con 🛠️

* [VueJs](https://vuejs.org/) - Framework web utilizado
* [Tailwind CSS](https://maven.apache.org/) - Framework CSS
* [ZincSearch](https://openobserve.ai/docs/quickstart/#openobserve-cloud) - Base de datos
* [GO](https://go.dev/) - Lenguaje utilizado en el indexer y el API
* [chi](https://go.dev/) - API Router
* [Axios](https://axios-http.com/docs/intro) - Cliente HTTP

## Autor ✒️

* **Oned Gómez** - *Trabajo Total* - [OnedGmez](https://github.com/OnedGmez)