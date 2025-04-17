export class Lead {
  id?: string
  email: string
  converted: boolean
  language: string

  constructor(
    email: string,
    converted: boolean,
    language: string,
    id?: string,
  ) {
    if (id) this.id = id
    this.email = email
    this.converted = converted
    this.language = language
  }
}
