type LeadParameters = {
  id?: string;
  email: string;
  converted: boolean;
  language?: string;
  createdAt?: Date;
  updatedAt?: Date;
};

export class Lead {
  id?: string;
  email: string;
  converted: boolean;
  language?: string;
  createdAt?: Date;
  updatedAt?: Date;

  constructor(parameters: LeadParameters) {
    this.id = parameters.id;
    this.email = parameters.email;
    this.converted = parameters.converted;
    this.language = parameters.language;
    this.createdAt = parameters.createdAt;
    this.updatedAt = parameters.updatedAt;
  }
}
