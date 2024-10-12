export interface IQueue {
  Consumer: (
    queueName: string,
    fn: (message: string) => Promise<void>,
  ) => Promise<void>
}
