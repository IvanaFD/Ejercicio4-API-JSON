# Ejercicio 4 - API JSON

## Stardew Villagers API

Simple REST API desarrollada en Go para gestionar la información de los aldeanos en Stardew Valley.  
La API permite consultar, agregar, actualizar y eliminar aldeanos utilizando un archivo JSON como almacenamiento.



## Puerto del servidor

El servidor se ejecuta en el puerto correspondiente al número de carnet:

```
http://localhost:24785
```



## Estructura del Proyecto

```
stardew-api/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── data/
│   └── villagers.json
│
├── internal/
│   ├── handlers/
|   |   └── info.go
|   |   └── villagers.go
|   |
│   ├── models/
|   |   └── villager.go
|   |
│   ├── storage/
|   |   └── json_store.go
|   |
│   └── utils/
|       └── error.go
|       └── http.go
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```
# Modelo de datos

Cada aldeano tiene la siguiente estructura:

```json
{
  "id": 1,
  "name": "Abigail",
  "birthday": "Fall 13",
  "marriageable": true,
  "best_gifts": ["Amethyst", "Chocolate Cake"],
  "location": "Pierre's General Store"
}
```


Campos:

| Campo | Tipo | Descripción |
|------|------|-------------|
| id | int | Identificador único |
| name | string | Nombre del aldeano |
| birthday | string | Cumpleaños dentro del juego |
| marriageable | bool | Indica si el personaje puede casarse |
| best_gifts | array | Lista de regalos favoritos |
| location | string | Lugar donde vive |

---


# Query Parameters

El endpoint `/api/villagers` permite utilizar filtros mediante parámetros en la URL.

| Parámetro | Tipo | Descripción |
|----------|------|-------------|
| id | int | Buscar aldeano por ID |
| location | string | Filtrar por ubicación |
| marriageable | bool | Filtrar por estado de matrimonio |

Los filtros pueden combinarse.

Ejemplo:

```
GET /api/villagers?location=2 Willow Lane&marriageable=true
```

---

# Ejemplos de uso

### Endpoint de información de la API

Este endpoint muestra información general sobre la API y los endpoints disponibles.

```
GET http://localhost:24785/
```

Ejemplo de respuesta:

```json
{
  "api": "Stardew Villagers API",
  "author": "Ivana FIgueroa",
  "description": "Simple REST API to manage Stardew Valley villagers",
  "endpoints": [
    "GET /api/villagers",
    "GET /api/villagers?id=1",
    "GET /api/villagers/{id}",
    "POST /api/villagers",
    "PUT /api/villagers?id=1",
    "DELETE /api/villagers?id=1"
  ],
  "version": "1.0"
}
```
---
### Obtener todos los aldeanos

```
GET http://localhost:24785/api/villagers
```

---

### Obtener aldeano por ID

```
GET http://localhost:24785/api/villagers?id=3
```

---

### Crear un nuevo aldeano

```
POST http://localhost:24785/api/villagers
```

Body ejemplo:

```json
{
  "name": "villager",
  "birthday": "Spring 1",
  "marriageable": true,
  "best_gifts": ["Coffee"],
  "location": "Farm"
}
```

---

### Actualizar un aldeano

```
PUT http://localhost:24785/api/villagers?id=3
```

---

### Eliminar un aldeano

```
DELETE http://localhost:24785/api/villagers?id=3
```



# Persistencia de datos

La API utiliza un archivo JSON para almacenar los datos:

```
data/villagers.json
```

Las operaciones:

- POST
- PUT
- DELETE

guardan automáticamente los cambios en este archivo.

---

# Ejecución del servidor

### Ejecutar localmente

```
go run cmd/server/main.go
```

Servidor disponible en:

```
http://localhost:24785
```

---

# Ejecutar con Docker

Construir y ejecutar el contenedor:

```
docker compose up --build
```

---

# Evidencia de funcionamiento

### 1. Servidor ejecutándose

El server corriendo de forma local con el endpoint que devuelve la información sobre la API

![Server corriendo](/images/cap1.png)

---

### 2. GET de todos los aldeanos

![GET todos los aldeanos](/images/cap2.png)

---

### 3. GET utilizando query parameters


![GET id 1](/images/cap3.png)
![GET id 1 con path](/images/cap4.png)
![GET filtros combinados](/images/cap5.png)


---

### 4. POST exitoso

![POST de un aldeano](/images/cap6.png)

---

### 5. PUT exitoso

se cambio la locación del aldeano que se creo en el PUT anterior
![Put de un aldeano](/images/cap7.png)

---

### 6. DELETE exitoso

Y finalmente se elimino el aldeno de la data
![DELETE del aldeano id 12](/images/cap8.png)


### 7. Casos de error

Respuesta al colocar id de un aldeano no registrado en la data
![id no encontrado](/images/cap9.png)

Respuesta al colocar mal los parametros
![id no encontrado](/images/cap10.png)


Respuesta al hacer un POST con un parámetro vacío
![id no encontrado](/images/cap11.png)

Respuesta al hacer un GET y no hay ningun aldeano que cumpla con los parámetros
![id no encontrado](/images/cap12.png)

# Tecnologías utilizadas

- Go (librería estándar)
- JSON como almacenamiento
- Docker