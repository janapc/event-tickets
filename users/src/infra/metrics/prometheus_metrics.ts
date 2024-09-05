import prometheus from 'prom-client'

const requestTotalCounter = new prometheus.Counter({
  name: 'request_total',
  help: 'Counter all requests',
  labelNames: ['route', 'method', 'statusCode'],
})

export { requestTotalCounter }
