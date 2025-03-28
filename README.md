# EjercicioAPI
API de Gestión de incidentes (tickets)

# API de Gestión de Incidentes

API para reportar y gestionar incidentes técnicos

# Requisitos
- Go 1.16+
- Echo

# Aclaración
Debido a que estoy usando tecnologías que no conocía tengo muchos archivos y carpetas siguiendo los tutoriales de dichas tecnologías, pero para probar el api sencilla con POST y GET se hace lo siguiente

## Instalación
1. Clonar repositorio
2. Entrar en la carpeta 'cmd' y luego en 'api'
3. Correr el comando 'go run main.go
4. Opcional: Se puede correr desde la carpeta principal cómo 'go run cmd/api/main.go'
5. Abril el siguiente enlace: http://localhost:1323
6. Probar get: 'curl http://localhost:1323/hello'
7. Probar POST: 'curl -X POST -H "Content-Type: application/json" -d '{"text": "Mensaje de prueba"}' http://localhost:1323/echo'

