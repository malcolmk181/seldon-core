# seldon-core-loadtesting

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square)

Loadtesting for Seldon Core v1

**Homepage:** <https://github.com/SeldonIO/seldon-core>

## Source Code

* <https://github.com/SeldonIO/seldon-core>
* <https://github.com/SeldonIO/seldon-core/tree/master/helm-charts/seldon-core-loadtesting>

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| data.size | int | `2` |  |
| image.release | float | `0.8` |  |
| loadtest.id | int | `1` |  |
| loadtest.sendFeedback | int | `0` |  |
| locust.clients | int | `10` |  |
| locust.hatchRate | int | `1` |  |
| locust.host | string | `"http://seldon-apiserver:8080"` |  |
| locust.maxWait | int | `1100` |  |
| locust.minWait | int | `990` |  |
| locust.script | string | `"predict_rest_locust.py"` |  |
| oauth.enabled | bool | `true` |  |
| oauth.key | string | `"key"` |  |
| oauth.secret | string | `"secret"` |  |
| replicaCount | int | `1` |  |
| rest.pathPrefix | string | `nil` |  |
