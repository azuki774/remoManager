# 使い方
    make build
    -> docker-compose up -d

## Metrics名
- remoMetrics_room_temperature
- remoMetrics_room_illuminance
- remoMetrics_room_moisture
- remoMetrics_room_human

# Prometheus側の設定

## URL
    http://<address>:9999/metrics
## Prometheus.yml
    - job_name: remo
    metrics_path: /metrics
    scrape_interval: 30s
    static_configs:
      - targets:
        - <address>:9999
