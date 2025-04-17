import prometheus from 'prom-client'

export function init(): void {
  const collectDefaultMetrics = prometheus.collectDefaultMetrics
  const register = prometheus.register
  collectDefaultMetrics({ register })
}
