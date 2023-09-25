This is a simple helm template that can use to turn up your akvorado namespace with the ipinfo container within your GKE cluster.

NOTE: The compatibility was tested only in GKE.
NOTE: The shared PV from the ipinfo was not tested.

Out of the box you will get the metrics of the Akvorado within your prometheus using the service discovery.

I'm also sharing a dashboard in Grafana that will provide a view of both the GKE Pods metrics and the Akvorado Metrics.

It's recomended the you create alerts using those metrics.

You can pass secrets directly in the helm template by placing within the `secrets` section inside the especific section. For `orchestartor` wold be `akvorado_orchestrator` and for `ipinfo` you can place it within `ipinfo`. This allow you to use enviroment variables to change the akvorado configuration or you can also update the configurations file inside the [config directory](config/).

You can also use a specific image of Akvorado in case you decided to run your own image and tag.

The configuration of the Akvorado itself should be done. Follows some of the parameters that you should configure:

- config/akvorado.yml
  - kafka: brokers
  - clickhouse: orchestrator_url
  - clickhouse: servers
  - username
  - password
  - database
  - asns

- config/console.yml: If you want to use redis for cache you should configure redis in this section

- config/inlet.yml
  - snmp: communities
  - classifiers


There's default limits for each deployments but you can override then, like the exemple:

```yaml
deployment:
  akvorado_inlet:
    resources:
      requests:
        memory: 800Mi
        cpu: 0.5
      limits:
        memory: 1000Mi
        cpu: 0.8
```

You can also change the args overriding the default( the default have only the application and the path to the configuration - e.g for inlet: inlet http://orchestrator.stage-akvorado.svc.cluster.local:8080)

```yaml
deployment:
  akvorado_inlet:
    args:
      - arg1
      - arg2
      - arg3
      - argn
```

Due the limition of the GKE in running only TCP or UDP in a service multiple services were created to the inlet:

- bmp: LoadBalancer service for the BMP that is configured for port 10179(not option was provided in the template to change this port)
- web: ClusterIP service for the inlet web services
- inlet: LoadBalancer service flow ingestion(default ports 2055 and 6343 not option is provided in the helm to user specific ports)

You can also change the volume sizes like the following exemple:

```yaml
volumes:
  akvorado_inlet:
    storageSize: 10Gi
```

The console web service(/) the inlet api(/api/v0/orchestrator) and the orchestrator api(/api/v0/orchestrator) were configured using a [nginx ingress service](templates/ingress-akvorado.yml) that's also have the `nginx.ingress.kubernetes.io/whitelist-source-range` that provides you a simple acl for the console front-end that you can also configure from the values file in the line 17 of the [values file](values-stage.yml)

The service for ingesting flows and bmp have a `loadBalancerSourceRanges` as well that you can configure the routers/switchs to receive flow in the line 20 of the exemple file.

There's other things that are also configurable using the template directly in the values file so I encorage you to check the templates and try it you.

This template does not provide a Kafka and a ClickHouse cluster for you.