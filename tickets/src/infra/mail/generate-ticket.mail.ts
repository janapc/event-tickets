interface TicketMailProps {
  name: string;
  passport: string;
  eventDescription: string;
  eventName: string;
  eventImageUrl: string;
  language: string;
}

export function generateTicketMail(props: TicketMailProps): {
  html: string;
  subject: string;
} {
  const body = formatBody(props.language);
  return {
    html: `
	<!DOCTYPE html>
<html lang="en">
<head>
	<style>
		body {
			display: flex;
			flex-direction: column;
			align-items: center;
		}
		img {
			width: 350px;
			height: 250px;
		}
		   .ticket {
				border: 2px dashed royalblue;
    		padding: 20px;
			 }
	</style>
</head>
<body>
	<h3>${props.eventName}</h3>
	<img src="${props.eventImageUrl}"/>
	<p>${props.eventDescription}</p>
	<p>${props.name} ${body.message}</p>
	<p class="ticket">${props.passport}</p>
</body>
</html>
	`,
    subject: body.subject,
  };
}

function formatBody(language: string): { message: string; subject: string } {
  switch (language) {
    case 'pt':
      return {
        subject: 'Seu ticket chegou =)',
        message: 'aqui está seu ticket:',
      };
    case 'en':
      return {
        subject: 'Your ticket is here =)',
        message: 'here is your passport:',
      };
    default:
      return {
        subject: 'Your ticket is here =)',
        message: 'here is your passport:',
      };
  }
}
