export function formatTemplate(
  name: string,
  passport: string,
  eventDescription: string,
  eventName: string,
  eventImageUrl: string,
): string {
  return `
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
	<h3>${eventName}</h3>
	<img src=${eventImageUrl}/>
	<p>${eventDescription}</p>
	<p>${name} here is your passport:</p>
	<p class="ticket">${passport}</p>
</body>
</html>
	`
}
