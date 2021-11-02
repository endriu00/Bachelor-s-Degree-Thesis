func (worker *warehouseBackendWorker) HandleMessage(message *nsq.Message) error {
  var userRequest Query
  var err error
  var wg sync.WaitGroup

  // Controllo sulla dimensione del corpo del messaggio ...
  
  // Timeout sul contesto
  ctx, cancel := context.WithTimeout(context.Background(), worker.deadline)
  defer cancel()
	
  // Decodifica del messaggio
  if err = json.Unmarshal(message.Body, &userRequest); err != nil {
  	worker.log.WithError(err).Error("Error unmarshaling nsq message body")
 	return err
  }

  // Codice relativo al log ...

  // Se non ci sono sismometri nella lista, si sono richiesti dati 
  // con una posizione altrimenti è stata fornita una lista di sismometri
  if len(userRequest.Seismometers) == 0 {
	if err = worker.GetResultsFromLocation(ctx, userRequest, &wg); 
	err == context.DeadlineExceeded {
	  // Codice su log ed errori ...
	}
	if err != nil {
	  // Codice su log ed errori ...
	}
  } else {
	  if err = worker.GetResultsFromDevices(ctx, userRequest, &wg); 
	  err == context.DeadlineExceeded {
	    // Codice su log ed errori ...
	  }
	  if err != nil {
	    // Codice su log ed errori ...
      }
  }
  wg.Wait()

  // Notifica dell'utente
  if err = worker.SendNotification(userRequest); err != nil {
 	// Codice su log ed errori ...
  }
  return nil
}