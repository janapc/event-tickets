export class Lead {
  id?: number
  email: string
  converted: boolean
  language: string

  constructor(
    email: string,
    converted: boolean,
    language: string,
    id?: number,
  ) {
    if (id) this.id = id
    this.email = email
    this.converted = converted
    this.language = language
  }
}
