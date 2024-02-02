# Weather Dashboard using Grafana

## Overview

This project implements a weather dashboard using Grafana, providing visualizations and insights into weather data. The architecture and forecast flow are illustrated below.

### Architecture

![Architecture](/images/WeatherArchitecture.svg)

### Forecast Flow

![ForecastFlow](/images/WeatherForecastFlow.svg)

## Usage

To get started, follow the instructions below:

1. Clone the repository:

   ```bash
   git clone https://github.com/shostyi-ve/apka.git
   cd apka
   ```

2. Build the Docker containers:

    ```bash
    make build
    ```

3. Launch the application:

    ```bash
    make up
    ```

    The application can now be accessed at http://localhost:3000

# Makefile Commands

- **build**: Builds the Docker containers.
- **up**: Launches the application in detached mode.
- **down**: Stops and removes the running containers.
- **clean**: Prunes Docker volumes forcefully.
- **restart**: Combines down, clean, build, and up commands for a quick restart.
