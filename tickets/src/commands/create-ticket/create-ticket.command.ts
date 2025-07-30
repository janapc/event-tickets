export interface CreateTicketCommandProps {
  messageId: string;
  name: string;
  email: string;
  eventId: string;
  eventName: string;
  eventDescription: string;
  eventImageUrl: string;
  language: string;
}

export class CreateTicketCommand {
  messageId: string;
  name: string;
  email: string;
  eventId: string;
  eventName: string;
  eventDescription: string;
  eventImageUrl: string;
  language: string;

  constructor(props: CreateTicketCommandProps) {
    this.messageId = props.messageId;
    this.name = props.name;
    this.email = props.email;
    this.eventId = props.eventId;
    this.eventName = props.eventName;
    this.eventDescription = props.eventDescription;
    this.eventImageUrl = props.eventImageUrl;
    this.language = props.language;
  }
}
