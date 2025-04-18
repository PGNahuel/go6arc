# GO 6 Arc
Este repositorio tiene una estructura simple de arquitectura hexagonal.

## Estructura de carpetas

```
/myapp
├── cmd/
│   ├── api/                # Punto de entrada principal de la API
│   │   ├── main.go         # Arranque del servidor HTTP
│   ├── worker/             # Opcional: si tienes procesos en segundo plano
│   │   ├── main.go         # Arranque de workers (ej. procesadores de eventos)
├── config/                 # Configuración de la app (archivos YAML, JSON, ENV)
│   ├── config.yaml
├── internal/               # Lógica interna (no se expone fuera del módulo)
│   ├── usecases/           # Casos de uso (application service layer)
│   │   ├── service.go      # Servicios de aplicación
│   ├── domain/             # Entidades de negocio (Enterprise Business Rules)
│   │   ├── user.go         # Ejemplo de entidad
│   │   ├── order.go
│   ├── ports/              # Interfaces (puertos de entrada/salida)
│   │   ├── http.go         # Interface para el adaptador HTTP
│   │   ├── repository.go   # Interface para persistencia
│   │   ├── external.go     # Interface para APIs externas
│   ├── adapters/           # Implementaciones de interfaces (adaptadores)
│   │   ├── http/           # Adaptador HTTP (manejo de requests)
│   │   │   ├── handler.go
│   │   ├── mysql/          # Adaptador de persistencia en MySQL
│   │   │   ├── repository.go
│   │   ├── external/       # Adaptador de APIs externas
│   │   │   ├── client.go
├── pkg/                    # Librerías reutilizables (helpers, utils)
│   ├── logger/             # Logger centralizado
│   │   ├── logger.go
│   ├── database/           # Conexión a MySQL
│   │   ├── mysql.go
│   ├── httpclient/         # Cliente HTTP genérico
│   │   ├── client.go
├── migrations/             # Archivos de migración de la base de datos
│   ├── 001_create_users.sql
│   ├── 002_create_orders.sql
├── test/                   # Pruebas unitarias e integración
│   ├── mock/               # Mocks de dependencias
│   ├── integration/        # Pruebas de integración
│   ├── unit/               # Pruebas unitarias
├── Makefile                # Comandos de build, test, etc.
├── go.mod                  # Módulo Go
├── go.sum                  # Dependencias
├── README.md               # Documentación del proyecto
```