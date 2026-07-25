# Project Structure

This project follows a layered architecture to keep the codebase clean, modular, and easy to maintain. Each directory has a specific responsibility.

```
uber-platform/

│
├── services/
│   │
│   ├── gateway-go/
│   │
│   ├── ride-service-go/
│   │
│   ├── payment-service-go/
│   │
│   ├── location-service-go/
│   │
│   ├── auth-service-node/
│   │
│   ├── user-service-node/
│   │
│   ├── notification-service-node/
│   │
│   └── analytics-service-node/
│
├── proto/
│
├── docker/
│
├── deployments/
│      ├── docker-compose.yml
│      └── kubernetes/
│
├── docs/
│
├── scripts/
│
└── README.md
```
