func (worker *_warehouseWorker) GetResultsFromDevices(ctx context.Context, 
  userRequest Query, wg *sync.WaitGroup) error {
  var err error
  pipeReader, pipeWriter := io.Pipe()
  writer := zip.NewWriter(pipeWriter)

  // Chiusura delle risorse con defer ...

  // Operazioni preeliminari ... 

  // Funzione asincrona per il consumo dei dati nel lato reader della Pipe.
  wg.Add(1)
  go func() {
   defer wg.Done()
   contentType := "application/zip"
   if err = worker.fs.PutWarehouseResults(ctx, userRequest.ID+".zip", 
   pipeReader, -1, minio.PutObjectOptions{ContentType: contentType}); err != nil {
	    // Log ed errori ... 
   }
  }()

  // Passi da compiere per ogni attributo richiesto.
  if userRequest.Attributes.Temperature {
	rows, err := worker.GetTemperatureQuery(ctx, userRequest, queryArgs)
	if err != nil {
		// Log ed errori ... 

	}
	if err = worker.CheckContextStatus(ctx); err != nil {
	    // Log ed errori ... 

	}
	if err = worker.UploadQueryResultsByDevices(userRequest, rows, 
	writer, "temperature"); err != nil && err != ErrNoRows {
		// Log ed errori ... 

	}
	if err = worker.CheckContextStatus(ctx); err != nil {
	    // Log ed errori ... 
	}
  }
  if userRequest.Attributes.RSSI {
    // Come per Temperature
  }
  if userRequest.Attributes.Quake {
    // Come per Temperature
  }
  if userRequest.Attributes.Threshold {
    // Come per Temperature
  }
  return nil
}