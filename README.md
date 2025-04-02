# API de Incidentes con Golang, MongoDB y Gorilla Mux

Este proyecto es una API REST que permite gestionar reportes de incidentes (computadoras, impresoras, redes, etc.) mediante operaciones CRUD. La API está desarrollada en Go, utiliza MongoDB para el almacenamiento y Gorilla Mux para el ruteo de las solicitudes HTTP.

## Índice

- [Requisitos Previos](#requisitos-previos)
- [Instalación y Configuración](#instalación-y-configuración)
  - [1. Instalar Go](#1-instalar-go)
  - [2. Instalar MongoDB](#2-instalar-mongodb)
  - [3. Clonar el Repositorio](#3-clonar-el-repositorio)
  - [4. Configurar el Proyecto](#4-configurar-el-proyecto)
- [Ejecutar la API](#ejecutar-la-api)
- [Insertar Datos de Prueba en MongoDB](#insertar-datos-de-prueba-en-mongodb)
- [Ejemplos de Uso](#ejemplos-de-uso)
- [Recursos Adicionales](#recursos-adicionales)

## Requisitos Previos

Antes de comenzar, asegúrate de tener instalados los siguientes componentes:

1. **Go**: Lenguaje de programación.
2. **MongoDB**: Base de datos NoSQL (puedes usar MongoDB Atlas o instalarlo localmente).
3. **Git**: Para clonar el repositorio.

## Instalación y Configuración

### 1. Instalar Go

- Descarga e instala Go desde [aquí](https://golang.org/dl/).
- Verifica la instalación con:
  ```bash
  go version
  ```

### 2. Instalar MongoDB

#### Opción A: MongoDB Atlas
- Regístrate en [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) y crea un clúster.
- Configura el acceso y usa la siguiente URI de conexión:
  ```
  mongodb+srv://JFMG:contraseña123@jfmg.xzezjjy.mongodb.net/?retryWrites=true&w=majority&appName=JFMG
  ```

#### Opción B: Instalación Local
- Descarga MongoDB desde [aquí](https://www.mongodb.com/try/download/community) y sigue las instrucciones de instalación.

### 3. Clonar el Repositorio

```bash
git clone https://github.com/MENENDEZGUERRA/EjercicioAPI.git
cd EjercicioAPI
cd cmd
cd api
```

### 4. Configurar el Proyecto

1. Inicializa el módulo de Go:
   ```bash
   go mod init EjercicioAPI
   ```
2. Instala las dependencias:
   ```bash
   go get -u github.com/gorilla/mux
   go get go.mongodb.org/mongo-driver/mongo
   ```

## Ejecutar la API

```bash
go run main.go
```

La API estará disponible en `http://localhost:8000`.

## Recomendación:
Utilizar la página https://www.postman.com instalando su cliente local para poder utilizar el lucalhost y de esa manera facilitar el uso de 'GET' 'POST' 'PUT' 'DELETE'

## Insertar Datos de Prueba en MongoDB

```js
db.incidents.insertMany([
  { reporter: "Juan Perez", description: "La computadora no enciende.", status: "pendiente", created_at: new Date() },
  { reporter: "Maria Lopez", description: "La impresora imprime mal.", status: "en proceso", created_at: new Date() },
  { reporter: "Carlos Sanchez", description: "La red es lenta.", status: "resuelto", created_at: new Date() }
])
```

## Ejemplos de Uso

### Crear un Incidente
```json
{
    "reporter": "Juan Perez",
    "description": "La computadora no enciende."
}
```

### Obtener Todos los Incidentes
```bash
GET /incidents
```

### Obtener un Incidente Específico
```bash
GET /incidents/{id}
```

### Actualizar el Estado de un Incidente
```json
{
    "status": "en proceso"
}
```

### Eliminar un Incidente
```bash
DELETE /incidents/{id}
```
 