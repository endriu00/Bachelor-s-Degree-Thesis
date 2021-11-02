func (worker *_warehouseWorker) UploadCSVByLocation(userRequest Query, 
    rows pgx.Rows, zipWriter *zip.Writer, fileName string) error {
  var csvWriter *csv.Writer
  var r Result

  fileInZipCSV, err := zipWriter.Create(fileName + ".csv")
  if err != nil {
	// Errori e log ...
  }
  csvWriter = csv.NewWriter(fileInZipCSV)

  if fileName == "quake" {
    // Scrittura dei campi che si andranno a scrivere ...
    // Gestione casistiche ... 
	for rows.Next() {
	  if err = rows.Scan(&r.Time, &r.SeismoId, &r.Quake, 
	  &r.DeviceLatitude, &r.DeviceLongitude); err != nil {
	      // Errori e log ...
	  }
	  // Operazioni per richiesta Location
	  csvLine := []string{r.SeismoId, r.Time.String(), r.Quake.String()}
	  if err = csvWriter.Write(csvLine); err != nil {
	      // Errori e log ...
	  }
	}
  }
  if fileName == "temperature" {
    // Come sopra ... 
  }
  if fileName == "rssi" {
    // Come sopra ... 
  }
  if fileName == "threshold" {
	// Come sopra ... 
  }
  csvWriter.Flush()
  return nil
}
