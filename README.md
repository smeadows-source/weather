# This package provides a very simple Golang implementation of a weather service.

## Endpoint
http://localhost:8080/weather

## Request
```
{
    "latitude": "40.754864",
    "longitude": "-74.007156"
}
```
## Response
```
{
    "Time": "This Afternoon",
    "summary": "Partly Sunny",
    "details": "Partly sunny. High near 83, with temperatures falling to around 80 in the afternoon. Northeast wind around 7 mph."
}
