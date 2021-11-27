package game

// Game describes a complete game content (graphics, controller, etc...)
type Game interface {
	Name() string
	// InitializeBuckets allow loading game components into engine
	InitializeBuckets() []EntityBucket
	// InitializeGame lets you load 1st scene (enable entity)
	InitializeGame(engine *Engine)
}
