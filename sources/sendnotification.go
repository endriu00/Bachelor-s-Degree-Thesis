func (worker *_warehouseWorker) SendNotification(userRequest Query) error {
  var err error
  if userRequest.Mail != "" {
    // Creazione del messaggio ...
    if err = worker.SendEmail([]string{userRequest.Mail}, emailMessage); 
    err != nil {
      // Log ed errori ...
    }
  }
  if userRequest.Tg != "" {
    // Creazione del messaggio ...
    if err = worker.SendTelegramMessage(userRequest.Tg, telegramMessage);
    err != nil {
      // Log ed errori ...
    }
  }
  return nil
}