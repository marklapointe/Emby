package repository

import (
	"database/sql"
	"fmt"
	"time"
)

// ItemRepository handles media item storage and retrieval.
type ItemRepository struct {
	*BaseRepository
}

// NewItemRepository creates a new item repository.
func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{BaseRepository: NewBaseRepository(db)}
}

// CreateSchema creates the necessary database tables if they don't exist.
func (r *ItemRepository) CreateSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS Items (
		Id TEXT PRIMARY KEY,
		Name TEXT NOT NULL,
		Overview TEXT,
		Tagline TEXT,
		IndexNumber INTEGER,
		ParentIndex INTEGER,
		CommunityRating REAL,
		RunTimeTicks INTEGER,
		ProductionYear INTEGER,
		OfficialRating TEXT,
		ContentType TEXT,
		MediaType TEXT,
		Genres TEXT,
		Studios TEXT,
		SeasonNumber INTEGER,
		EpisodeNumber INTEGER,
		Album TEXT,
		Artists TEXT,
		ExtraType TEXT,
		ChannelNumber INTEGER,
		StartDate TEXT,
		EndDate TEXT,
		IsLive INTEGER,
		IsSeries INTEGER,
		IsMovie INTEGER,
		IsNews INTEGER,
		IsSports INTEGER,
		IsKids INTEGER,
		IsPremiere INTEGER,
		LocationType TEXT,
		Path TEXT,
		PrimaryImageURL TEXT,
		BackdropImageURL TEXT,
		ParentID TEXT,
		Width INTEGER,
		Height INTEGER,
		Video3DFormat TEXT,
		PostLiveFeedTime INTEGER,
		LiveMediaSourceID TEXT,
		StartTimeTicks INTEGER,
		EndTimeTicks INTEGER,
		RemoteImageURL TEXT,
		LocalTrailerCount INTEGER,
		LockedFields TEXT,
		LockData INTEGER,
		Disabled INTEGER,
		EnableMediaSourceDisplay INTEGER,
		ExtraIds TEXT,
		CreatedDate TEXT,
		ModifiedDate TEXT
	);

	CREATE TABLE IF NOT EXISTS MediaSources (
		Id TEXT PRIMARY KEY,
		ItemId TEXT NOT NULL,
		Name TEXT,
		Type TEXT,
		Container TEXT,
		Size INTEGER,
		Path TEXT,
		Protocol TEXT,
		Encoder INTEGER,
		VideoCodec TEXT,
		AudioCodec TEXT,
		Format TEXT,
		Width INTEGER,
		Height INTEGER,
		RefFrames INTEGER,
		VideoFramerate TEXT,
		VideoBitRate INTEGER,
		AudioBitRate INTEGER,
		AudioChannels INTEGER,
		AudioSampleRate TEXT,
		DefaultAudioStreamIndex INTEGER,
		SupportsTranscoding INTEGER,
		SupportsDirectStream INTEGER,
		SupportsDirectPlay INTEGER,
		IsRemote INTEGER,
		FOREIGN KEY(ItemId) REFERENCES Items(Id)
	);

	CREATE TABLE IF NOT EXISTS Users (
		Id TEXT PRIMARY KEY,
		Name TEXT NOT NULL,
		Username TEXT,
		EmailAddress TEXT,
		LoginUsername TEXT,
		LoginPassword TEXT,
		InvalidLoginAttemptCount INTEGER,
		LastLoginDate TEXT,
		LastActivityDate TEXT,
		AuthenticationProviderID TEXT,
		PrimaryImageTag TEXT,
		Policy TEXT
	);

	CREATE TABLE IF NOT EXISTS UserItems (
		Id TEXT PRIMARY KEY,
		UserId TEXT NOT NULL,
		ItemID TEXT NOT NULL,
		PlaybackPositionTicks INTEGER,
		PlayCount INTEGER,
		IsFavorite INTEGER,
		Liked INTEGER,
		LastPlayedDate TEXT,
		Played INTEGER,
		Rating INTEGER,
		FOREIGN KEY(UserId) REFERENCES Users(Id),
		FOREIGN KEY(ItemID) REFERENCES Items(Id)
	);

	CREATE TABLE IF NOT EXISTS Sessions (
		Id TEXT PRIMARY KEY,
		Client TEXT,
		DeviceName TEXT,
		DisplayName TEXT,
		Endpoint TEXT,
		LocalAddress TEXT,
		RemoteAddress TEXT,
		MachineId TEXT,
		LastActivityTime TEXT,
		LastPlaybackTime TEXT,
		PlaybackPositionTicks INTEGER,
		PlayMethod TEXT,
		SupportsMediaControl INTEGER,
		SupportsPersistentIdentification INTEGER,
		SupportsSync INTEGER,
		IsInActiveSession INTEGER,
		IsTerminal INTEGER,
		StartTimeTicks INTEGER
	);

	CREATE INDEX IF NOT EXISTS idx_items_path ON Items(Path);
	CREATE INDEX IF NOT EXISTS idx_items_mediatype ON Items(MediaType);
	CREATE INDEX IF NOT EXISTS idx_items_locationtype ON Items(LocationType);
	CREATE INDEX IF NOT EXISTS idx_items_parentid ON Items(ParentID);
	CREATE INDEX IF NOT EXISTS idx_mediasources_itemid ON MediaSources(ItemId);
	CREATE INDEX IF NOT EXISTS idx_useritems_userid ON UserItems(UserId);
	CREATE INDEX IF NOT EXISTS idx_useritems_itemid ON UserItems(ItemID);
	`

	_, err := r.Exec(schema)
	return err
}

// InsertItem inserts a media item into the database.
func (r *ItemRepository) InsertItem(id, name, path, locationType string) error {
	query := `
		INSERT OR REPLACE INTO Items 
		(Id, Name, Path, LocationType, ContentType, CreatedDate, ModifiedDate)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now().UTC().Format(time.RFC3339)
	_, err := r.Exec(query, id, name, path, locationType, locationType, now, now)
	return err
}

// GetItem retrieves a media item by ID.
func (r *ItemRepository) GetItem(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, Tagline, ContentType, MediaType, 
		       Path, LocationType, ProductionYear, CommunityRating,
		       RunTimeTicks, IsMovie, IsSeries, IsLive
		FROM Items WHERE Id = ?
	`
	row := r.QueryRow(query, id)
	
	var item map[string]interface{} = make(map[string]interface{})
	var isMovie, isSeries, isLive int
	var overview, tagline, contentType, mediaType, path, locationType sql.NullString
	var productionYear sql.NullInt64
	var communityRating sql.NullFloat64
	var runTimeTicks sql.NullInt64
	
	var idVal, nameVal string
	err := row.Scan(&idVal, &nameVal, &overview, &tagline,
		&contentType, &mediaType, &path, &locationType,
		&productionYear, &communityRating, &runTimeTicks,
		&isMovie, &isSeries, &isLive,
	)
	if err != nil {
		return nil, fmt.Errorf("scan item: %w", err)
	}

	item["Id"] = idVal
	item["Name"] = nameVal
	if overview.Valid {
		item["Overview"] = overview.String
	}
	if tagline.Valid {
		item["Tagline"] = tagline.String
	}
	if contentType.Valid {
		item["ContentType"] = contentType.String
	}
	if mediaType.Valid {
		item["MediaType"] = mediaType.String
	}
	if path.Valid {
		item["Path"] = path.String
	}
	if locationType.Valid {
		item["LocationType"] = locationType.String
	}
	if productionYear.Valid {
		item["ProductionYear"] = int(productionYear.Int64)
	}
	if communityRating.Valid {
		item["CommunityRating"] = communityRating.Float64
	}
	if runTimeTicks.Valid {
		item["RunTimeTicks"] = runTimeTicks.Int64
	}
	item["IsMovie"] = isMovie == 1
	item["IsSeries"] = isSeries == 1
	item["IsLive"] = isLive == 1

	return item, nil
}

// SearchItems searches for media items by name.
func (r *ItemRepository) SearchItems(query string, limit, offset int) ([]map[string]interface{}, error) {
	sqlQuery := `
		SELECT Id, Name, ContentType, MediaType, Path, 
		       ProductionYear, CommunityRating, IsMovie, IsSeries
		FROM Items 
		WHERE Name LIKE ? 
		ORDER BY Name 
		LIMIT ? OFFSET ?
	`
	rows, err := r.Query(sqlQuery, "%"+query+"%", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query items: %w", err)
	}
	defer rows.Close()

	var items []map[string]interface{}
	for rows.Next() {
		var idVal, nameVal string
		var contentType, mediaType, path sql.NullString
		var productionYear sql.NullInt64
		var communityRating sql.NullFloat64
		var isMovie, isSeries sql.NullInt64
		
		err := rows.Scan(&idVal, &nameVal, &contentType, &mediaType,
			&path, &productionYear, &communityRating,
			&isMovie, &isSeries,
		)
		if err != nil {
			return nil, fmt.Errorf("scan item: %w", err)
		}

		item := make(map[string]interface{})
		item["Id"] = idVal
		item["Name"] = nameVal
		if contentType.Valid {
			item["ContentType"] = contentType.String
		}
		if mediaType.Valid {
			item["MediaType"] = mediaType.String
		}
		if path.Valid {
			item["Path"] = path.String
		}
		if productionYear.Valid {
			item["ProductionYear"] = int(productionYear.Int64)
		}
		if communityRating.Valid {
			item["CommunityRating"] = communityRating.Float64
		}
		if isMovie.Valid {
			item["IsMovie"] = isMovie.Int64 == 1
		}
		if isSeries.Valid {
			item["IsSeries"] = isSeries.Int64 == 1
		}

		items = append(items, item)
	}
	
	return items, rows.Err()
}

// GetTotalItemCounts returns counts by content type.
func (r *ItemRepository) GetTotalItemCounts() (map[string]int, error) {
	query := `
		SELECT ContentType, COUNT(*) as Count 
		FROM Items 
		GROUP BY ContentType
	`
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)
	for rows.Next() {
		var contentType sql.NullString
		var count int
		if err := rows.Scan(&contentType, &count); err != nil {
			return nil, err
		}
		if contentType.Valid {
			counts[contentType.String] = count
		}
	}
	
	return counts, rows.Err()
}
