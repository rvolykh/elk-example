# ELK stack example

Simple example hot to ship your data to Elasticsearch.
- filebeat (load csv files to ELK)
- logrus (logs to ELK)

### Requirements

- docker
- docker-compose
- go
- python

### Workflow
1. docker-compose up -d
2. check http://localhost:9200
3. cd filebeat-5.4.0-linux-x86_64
4. ./filebeat -e
5. check http://localhost:9200/elk-example/_search
6. cd ..
7. python kibana_io.py import --dir ./kibana --url http://localhost:9200
8. check http://localhost:5601
9. select elk-example as default index
10. explore dashboards
11. go run logrus-elk-hook.go
12. check http://localhost:9200/elk-example/json/_search

