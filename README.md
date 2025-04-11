# Ultra-Fast Go Ad Services
Server A high-performance, Go-based microservices architecture designed for seamless ad service delivery. This repository supports robust advertising platforms, including DSP, SSP, SSAI, and more, ensuring scalable and efficient ad operations.


/go-ad-service
├── /api-gateway       # Entry point for all external traffic
├── /bidder            # Core bidding engine microservice
├── /inventory-manager # Supply inventory processing
├── /user-data         # User profile and targeting data
├── /campaign-manager  # Campaign configuration and management
├── /analytics         # Real-time and historical analytics
├── /shared            # Shared libraries and utilities
│   ├── /proto         # Protocol buffer definitions
│   ├── /metrics       # Metrics collection utilities
│   └── /models        # Shared data models
└── /deploy            # Deployment configurations