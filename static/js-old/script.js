function changeMessage() {
  const messageElement =document.getElementById('message');
  const currentMessage = messageElement.textContent;

  const alternateMessage =[
    "Ani",
    "Nadine",
    "Babitida",
    "Hortencia",
    "Diana",
    "Angie",
    "mariana",
    "Ndo",
  ]

  let newMessage;

  do {
    newMessage = alternateMessage[Math.floor(Math.random() * alternateMessage.length)];
  } while (newMessage === currentMessage);

  messageElement.textContent = newMessage;

}