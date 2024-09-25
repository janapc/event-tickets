import prometheus from 'prom-client'

export function init() {
  const collectDefaultMetrics = prometheus.collectDefaultMetrics
  const register = prometheus.register
  collectDefaultMetrics({ register })
}
