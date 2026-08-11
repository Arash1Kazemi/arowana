package storage

// BaseRepository provides shared scaffolding that concrete repositories
// embed to avoid repeating connection-handling code.
type BaseRepository struct {
	table string
}

func (r BaseRepository) TableName() string {
	return r.table
}

// UserRepository embeds BaseRepository, so it inherits TableName() for
// free — this is the relationship the embeds edge kind needs to capture.
type UserRepository struct {
	BaseRepository
	CacheEnabled bool
}