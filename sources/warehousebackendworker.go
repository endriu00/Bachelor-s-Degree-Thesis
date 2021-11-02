type warehouseBackendWorker struct {
	fs               storage.SCSFileStorage
	timescaleDBPool  *pgx.Pool
	telegramBotToken string
	telegramApiUrl   string
	db               *sqlx.DB
	smtpConfig       SMTPConfig
	deadline         time.Duration
	ctx              context.Context
	log              *logrus.Entry
	// elementi per la concorrenza
}