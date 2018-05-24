# TP 1: Proxy - Tecno: Golang

## Arquitectura

- main.go: archivo central del proxy, se encarga de concentrar su funcionalidad. Recibe peticiones del browser, loguea url junto con la fecha (gorutine), realiza un manejo de cache con ccache para saber si ya se estuvo en la pagina pedida (gorutine), se realiza la petición de la página solicitada y se devuelve la información al browser( con algunas modificaciones realizadas por el  mwHtml).
- loger.go: es el archivo encargado de todas las registraciones de logs. Mantiene  el archivo ¨log-file¨ con informacion de Fecha y url de la petición. Tambien registra, en la carpeta ¨request-logs¨, el body de todos los archivos pedidos.
- mw-html.go: se encarga de inyectar bytes adicionales a lo enviado como respuestas al browser.

## Configuracion del browser

En el browser ir a Settings > advanced > System > Open Proxy settings.

![Browser setting](https://i.imgur.com/PiMDTuU.png)

Ip: localhost
defaultPort: 8000

## Como Correr la aplicación ?

```bash
git clone repo
cd repo
go run main.go [portNum]
or
go build main.go
./main [portNum]
```

## Comparación Runtime

Runtime sin gorutines
![It doesn't have gorutine](https://i.imgur.com/JYtYErU.png)

Runtime con gorutines
![It has gorutine](https://i.imgur.com/yNRUbMy.png)
