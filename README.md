Nature Remo Cloud API を用いて、Prometheus上にデータを展開するプログラム。
Nature Remo Cloud APIのアクセスリミットに対応し、1分間に1度だけデータを更新する。

# 使い方
    make build
    -> docker-compose up -d

で起動する。自動的に1分間に1回データを取得しに行き、

    http://<address>:8800/sensor

で取得済のデータを取得する。

- Prometheusで利用する場合は、 `/build` ディレクトリに `token.env` を配置する。
- 中身は `APPLIANCE_ID`（照明用、not yet implemented）と、 `TOKEN` （Remo API アクセストークン）を記載する。

## PrometheusのMetrics名
- remoMetrics_room_temperature
- remoMetrics_room_illuminance
- remoMetrics_room_motion
- remoMetrics_room_humidity

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
