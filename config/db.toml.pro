[rabbitmq_ichunt]
queue_name="ichunt_monitor_user_behavior"
routing_key="ichunt_monitor_user_behavior"
exchange="ichunt_monitor_behavior"
type="direct"
dns="amqp://huntmouser:jy2y2900@172.18.137.23:5672/"

[mongodb_databases]
dns="mongodb://Tzmonitor:husttmon6tpz99@172.18.137.23:27017/monitor?authMechanism=SCRAM-SHA-1"
databases="monitor"
collection="monitor_log"

[logSaveConfig]
time=120
length=1000