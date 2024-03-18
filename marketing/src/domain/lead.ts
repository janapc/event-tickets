export class Lead {
  id?: number
  email: string
  converted: boolean

  constructor(email: string, converted: boolean, id?: number) {
    if (id) this.id = id
    this.email = email
    this.converted = converted
  }
}
