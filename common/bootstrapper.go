package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initalize private/public keys for JWT authentication
	initKeys()
	// Start a MongoDb session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
