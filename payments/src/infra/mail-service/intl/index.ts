export const IntlApprovedPayment = (
  language: string,
  name: string,
): { subject: string; message: string } => {
  switch (language) {
    case 'pt':
      return {
        subject: 'Pagamento Aprovado =)',
        message: `Olá ${name}.\n Seu pagamento foi aprovado.`,
      };
    case 'en':
      return {
        subject: 'Payment Approved =)',
        message: `Hello ${name}.\n Your payment has been approved.`,
      };
    default:
      break;
  }
};

export const IntlRejectedPayment = (
  language: string,
  name: string,
): { subject: string; message: string } => {
  switch (language) {
    case 'pt':
      return {
        subject: 'Pagamento Rejeitado =(',
        message: `Olá ${name}.\n Seu pagamento foi rejeitado, porfavor tente novamente no site.`,
      };
    case 'en':
      return {
        subject: 'Payment Rejected =(',
        message: `Hello ${name}.\n Your payment has been rejected, try again in us site`,
      };
    default:
      break;
  }
};
