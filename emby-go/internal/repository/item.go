package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/emby/emby-go/internal/model"
	"gorm.io/gorm"
)

// ItemRepository handles media item storage and retrieval.
type ItemRepository struct {
	*BaseRepository
}

// NewItemRepository creates a new item repository.
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{BaseRepository: NewBaseRepository(db)}
}

// CreateSchema creates the necessary database tables if they don't exist using GORM AutoMigrate.
func (r *ItemRepository) CreateSchema() error {
	return r.db.AutoMigrate(
		&model.GORMItem{},
		&model.GORMItemMediaType{},
		&model.GORMMediaSource{},
		&model.GORMUser{},
		&model.GORMUserItem{},
		&model.GORMSession{},
	)
}

// GetItemsByParent returns items with a given parent ID.
func (r *ItemRepository) GetItemsByParent(parentID string, mediaType string, limit, offset int) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, ContentType, MediaType, Path,
		       ProductionYear, CommunityRating,
		       RunTimeTicks, PrimaryImageURL
		FROM Items
		WHERE ParentID = ?
	`
	args := []interface{}{parentID}

	if mediaType != "" {
		query += " AND MediaType = ?"
		args = append(args, mediaType)
	}

	query += " ORDER BY Name LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query items by parent: %w", err)
	}
	defer rows.Close()

	var items []map[string]interface{}
	for rows.Next() {
		var id, name, overview, contentType, mediaType, path, primaryImage sql.NullString
		var productionYear, runTimeTicks sql.NullInt64
		var communityRating sql.NullFloat64

		err := rows.Scan(&id, &name, &overview, &contentType, &mediaType, &path,
			&productionYear, &communityRating, &runTimeTicks, &primaryImage)
		if err != nil {
			return nil, fmt.Errorf("scan item: %w", err)
		}

		item := make(map[string]interface{})
		item["Id"] = id.String
		item["Name"] = name.String
		if overview.Valid {
			item["Overview"] = overview.String
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
		if productionYear.Valid {
			item["ProductionYear"] = int(productionYear.Int64)
		}
		if communityRating.Valid {
			item["CommunityRating"] = communityRating.Float64
		}
		if runTimeTicks.Valid {
			item["RunTimeTicks"] = runTimeTicks.Int64
		}
		if primaryImage.Valid {
			item["PrimaryImageUrl"] = primaryImage.String
		}

		items = append(items, item)
	}

	if items == nil {
		items = []map[string]interface{}{}
	}

	return items, rows.Err()
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
	if err != nil {
		return err
	}

	mediaType := mapLocationTypeToMediaType(locationType)
	if mediaType != "" {
		insertQuery := `INSERT INTO ItemMediaTypes (ItemId, MediaType) VALUES (?, ?)`
		_, err = r.Exec(insertQuery, id, mediaType)
		if err != nil {
			return err
		}
	}

	return nil
}

func mapLocationTypeToMediaType(locationType string) string {
	switch locationType {
	case "movies", "Movie", "Video":
		return "Movie"
	case "tvshows", "Series", "TvShows":
		return "Series"
	case "music", "Music", "Audio":
		return "Audio"
	case "photos", "Photos", "Photo":
		return "Photo"
	case "books", "Books", "Book":
		return "Book"
	case "games", "Games", "Game":
		return "Game"
	case "homevideos", "HomeVideos", "HomeVideo":
		return "HomeVideo"
	case "livetv", "LiveTV":
		return "LiveTV"
	default:
		if locationType != "" {
			return locationType
		}
		return ""
	}
}

// GetItem retrieves a media item by ID.
func (r *ItemRepository) GetItem(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, Tagline, ContentType, MediaType,
		       Path, LocationType, ProductionYear, CommunityRating,
		       RunTimeTicks
		FROM Items WHERE Id = ?
	`
	row := r.QueryRow(query, id)

	var item map[string]interface{} = make(map[string]interface{})
	var overview, tagline, contentType, mediaType, path, locationType sql.NullString
	var productionYear sql.NullInt64
	var communityRating sql.NullFloat64
	var runTimeTicks sql.NullInt64

	var idVal, nameVal string
	err := row.Scan(&idVal, &nameVal, &overview, &tagline,
		&contentType, &mediaType, &path, &locationType,
		&productionYear, &communityRating, &runTimeTicks,
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

	mediaTypes, err := r.GetItemMediaTypes(id)
	if err == nil {
		item["MediaTypes"] = mediaTypes
	}

	return item, nil
}

// GetAllItems returns all media items from the database.
func (r *ItemRepository) GetAllItems() ([]map[string]interface{}, error) {
	sqlQuery := `
		SELECT Id, Name, ContentType, MediaType, Path,
		       ProductionYear, CommunityRating
		FROM Items
		ORDER BY Name
	`

	rows, err := r.Query(sqlQuery)
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

		err := rows.Scan(&idVal, &nameVal, &contentType, &mediaType,
			&path, &productionYear, &communityRating,
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
			item["ProductionYear"] = productionYear.Int64
		}
		if communityRating.Valid {
			item["CommunityRating"] = communityRating.Float64
		}

		items = append(items, item)
	}

	if items == nil {
		items = []map[string]interface{}{}
	}

	return items, rows.Err()
}

// SearchItems searches for media items by name.
func (r *ItemRepository) SearchItems(query string, limit, offset int) ([]map[string]interface{}, error) {
	sqlQuery := `
		SELECT Id, Name, ContentType, MediaType, Path,
		       ProductionYear, CommunityRating
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

		err := rows.Scan(&idVal, &nameVal, &contentType, &mediaType,
			&path, &productionYear, &communityRating,
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

// GetChannels returns TV channels.
func (r *ItemRepository) GetChannels(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ChannelNumber, ExtraType as ChannelType,
		       PrimaryImageURL, BackdropImageURL, CreatedDate, ModifiedDate
		FROM Items
		WHERE MediaType = 'Channel' AND LocationType = 'Remote'
		ORDER BY ChannelNumber, Name`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var channels []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, channelType, primaryImage, backdropImage, createdDate, modifiedDate sql.NullString
		var channelNumber sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &channelNumber, &channelType,
			&primaryImage, &backdropImage, &createdDate, &modifiedDate); err != nil {
			return nil, err
		}
		
		channel := map[string]interface{}{
			"Id":           id.String,
			"Name":         name.String,
			"Overview":     overview.String,
			"MediaType":    mediaType.String,
			"ChannelNumber": channelNumber.Int64,
			"ChannelType":  channelType.String,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			channel["PrimaryImageUrl"] = primaryImage.String
		}
		if backdropImage.Valid && backdropImage.String != "" {
			channel["BackdropImageUrl"] = backdropImage.String
		}
		if createdDate.Valid {
			channel["CreateDate"] = createdDate.String
		}
		if modifiedDate.Valid {
			channel["DateModified"] = modifiedDate.String
		}
		
		channels = append(channels, channel)
	}
	
	if channels == nil {
		channels = []map[string]interface{}{}
	}
	
	return channels, rows.Err()
}

// GetChannel returns a single channel.
func (r *ItemRepository) GetChannel(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ChannelNumber, ExtraType as ChannelType,
		       PrimaryImageURL, BackdropImageURL, CreatedDate, ModifiedDate
		FROM Items
		WHERE Id = ? AND MediaType = 'Channel'`
	
	var itemId, name, overview, mediaType, channelType, primaryImage, backdropImage, createdDate, modifiedDate sql.NullString
	var channelNumber sql.NullInt64
	
	err := r.QueryRow(query, id).Scan(&itemId, &name, &overview, &mediaType, &channelNumber, &channelType,
		&primaryImage, &backdropImage, &createdDate, &modifiedDate)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("channel not found: %s", id)
	}
	if err != nil {
		return nil, err
	}
	
	channel := map[string]interface{}{
		"Id":            itemId.String,
		"Name":          name.String,
		"Overview":      overview.String,
		"MediaType":     mediaType.String,
		"ChannelNumber": channelNumber.Int64,
		"ChannelType":   channelType.String,
	}
	
	if primaryImage.Valid && primaryImage.String != "" {
		channel["PrimaryImageUrl"] = primaryImage.String
	}
	if backdropImage.Valid && backdropImage.String != "" {
		channel["BackdropImageUrl"] = backdropImage.String
	}
	if createdDate.Valid {
		channel["CreateDate"] = createdDate.String
	}
	if modifiedDate.Valid {
		channel["DateModified"] = modifiedDate.String
	}
	
	return channel, nil
}

// GetChannelFolders returns folders for a channel.
func (r *ItemRepository) GetChannelFolders(channelId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ExtraType as ContentType
		FROM Items
		WHERE ParentID = ? OR Id = ?
		ORDER BY Name`
	
	rows, err := r.Query(query, channelId, channelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var folders []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType sql.NullString
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType); err != nil {
			return nil, err
		}
		
		folders = append(folders, map[string]interface{}{
			"Id":          id.String,
			"Name":        name.String,
			"Overview":    overview.String,
			"MediaType":   mediaType.String,
			"ContentType": contentType.String,
			"Type":        "Folder",
		})
	}
	
	if folders == nil {
		folders = []map[string]interface{}{}
	}
	
	return folders, rows.Err()
}

// GetChannelItems returns items for a channel.
func (r *ItemRepository) GetChannelItems(channelId, userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType, 
		       RunTimeTicks, ProductionYear, OfficialRating,
		       PrimaryImageURL, BackdropImageURL
		FROM Items
		WHERE ParentID = ?
		ORDER BY SortName, Name`
	
	rows, err := r.Query(query, channelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType, officialRating, primaryImage, backdropImage sql.NullString
		var runTimeTicks, productionYear sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&runTimeTicks, &productionYear, &officialRating, &primaryImage, &backdropImage); err != nil {
			return nil, err
		}
		
		item := map[string]interface{}{
			"Id":              id.String,
			"Name":            name.String,
			"Overview":        overview.String,
			"MediaType":       mediaType.String,
			"Type":            contentType.String,
			"RunTimeTicks":    runTimeTicks.Int64,
			"ProductionYear":  productionYear.Int64,
			"OfficialRating":  officialRating.String,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			item["PrimaryImageUrl"] = primaryImage.String
		}
		if backdropImage.Valid && backdropImage.String != "" {
			item["BackdropImageUrl"] = backdropImage.String
		}
		
		items = append(items, item)
	}
	
	if items == nil {
		items = []map[string]interface{}{}
	}
	
	return items, rows.Err()
}

// GetPrograms returns TV programs.
func (r *ItemRepository) GetPrograms(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, EndTimeTicks, RunTimeTicks,
		       SeasonNumber, EpisodeNumber, ChannelNumber
		FROM Items
		WHERE MediaType = 'Video' AND ContentType = 'TvChannel'
		ORDER BY StartTimeTicks, Name`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var programs []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType sql.NullString
		var startTimeTicks, endTimeTicks, runTimeTicks, seasonNumber, episodeNumber, channelNumber sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&startTimeTicks, &endTimeTicks, &runTimeTicks, &seasonNumber, &episodeNumber, &channelNumber); err != nil {
			return nil, err
		}
		
		programs = append(programs, map[string]interface{}{
			"Id":              id.String,
			"Name":            name.String,
			"Overview":        overview.String,
			"MediaType":       mediaType.String,
			"Type":            contentType.String,
			"StartTimeTicks":  startTimeTicks.Int64,
			"EndTimeTicks":    endTimeTicks.Int64,
			"RunTimeTicks":    runTimeTicks.Int64,
			"SeasonNumber":    seasonNumber.Int64,
			"EpisodeNumber":   episodeNumber.Int64,
			"ChannelNumber":   channelNumber.Int64,
		})
	}
	
	if programs == nil {
		programs = []map[string]interface{}{}
	}
	
	return programs, rows.Err()
}

// GetProgram returns a single TV program.
func (r *ItemRepository) GetProgram(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, EndTimeTicks, RunTimeTicks,
		       SeasonNumber, EpisodeNumber, ChannelNumber
		FROM Items
		WHERE Id = ?`
	
	var itemId, name, overview, mediaType, contentType sql.NullString
	var startTimeTicks, endTimeTicks, runTimeTicks, seasonNumber, episodeNumber, channelNumber sql.NullInt64
	
	err := r.QueryRow(query, id).Scan(&itemId, &name, &overview, &mediaType, &contentType,
		&startTimeTicks, &endTimeTicks, &runTimeTicks, &seasonNumber, &episodeNumber, &channelNumber)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("program not found: %s", id)
	}
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"Id":              itemId.String,
		"Name":            name.String,
		"Overview":        overview.String,
		"MediaType":       mediaType.String,
		"Type":            contentType.String,
		"StartTimeTicks":  startTimeTicks.Int64,
		"EndTimeTicks":    endTimeTicks.Int64,
		"RunTimeTicks":    runTimeTicks.Int64,
		"SeasonNumber":    seasonNumber.Int64,
		"EpisodeNumber":   episodeNumber.Int64,
		"ChannelNumber":   channelNumber.Int64,
	}, nil
}

// GetRecordings returns TV recordings.
func (r *ItemRepository) GetRecordings(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, EndTimeTicks, RunTimeTicks,
		       ProductionYear, PrimaryImageURL
		FROM Items
		WHERE MediaType = 'Video' AND ContentType = 'Recording'
		ORDER BY StartTimeTicks DESC`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var recordings []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType, primaryImage sql.NullString
		var startTimeTicks, endTimeTicks, runTimeTicks, productionYear sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&startTimeTicks, &endTimeTicks, &runTimeTicks, &productionYear, &primaryImage); err != nil {
			return nil, err
		}
		
		recording := map[string]interface{}{
			"Id":              id.String,
			"Name":            name.String,
			"Overview":        overview.String,
			"MediaType":       mediaType.String,
			"Type":            contentType.String,
			"StartTimeTicks":  startTimeTicks.Int64,
			"EndTimeTicks":    endTimeTicks.Int64,
			"RunTimeTicks":    runTimeTicks.Int64,
			"ProductionYear":  productionYear.Int64,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			recording["PrimaryImageUrl"] = primaryImage.String
		}
		
		recordings = append(recordings, recording)
	}
	
	if recordings == nil {
		recordings = []map[string]interface{}{}
	}
	
	return recordings, rows.Err()
}

// GetRecording returns a single TV recording.
func (r *ItemRepository) GetRecording(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, EndTimeTicks, RunTimeTicks,
		       ProductionYear, PrimaryImageURL, Path
		FROM Items
		WHERE Id = ?`
	
	var itemId, name, overview, mediaType, contentType, primaryImage, path sql.NullString
	var startTimeTicks, endTimeTicks, runTimeTicks, productionYear sql.NullInt64
	
	err := r.QueryRow(query, id).Scan(&itemId, &name, &overview, &mediaType, &contentType,
		&startTimeTicks, &endTimeTicks, &runTimeTicks, &productionYear, &primaryImage, &path)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("recording not found: %s", id)
	}
	if err != nil {
		return nil, err
	}
	
	recording := map[string]interface{}{
		"Id":              itemId.String,
		"Name":            name.String,
		"Overview":        overview.String,
		"MediaType":       mediaType.String,
		"Type":            contentType.String,
		"StartTimeTicks":  startTimeTicks.Int64,
		"EndTimeTicks":    endTimeTicks.Int64,
		"RunTimeTicks":    runTimeTicks.Int64,
		"ProductionYear":  productionYear.Int64,
	}
	
	if primaryImage.Valid && primaryImage.String != "" {
		recording["PrimaryImageUrl"] = primaryImage.String
	}
	if path.Valid && path.String != "" {
		recording["Path"] = path.String
	}
	
	return recording, nil
}

// GetTimers returns TV timers.
func (r *ItemRepository) GetTimers(userId string) ([]map[string]interface{}, error) {
	// Return empty for now - timers would be in a separate table
	return []map[string]interface{}{}, nil
}

// GetChannelsWithImage returns TV channels with images.
func (r *ItemRepository) GetChannelsWithImage(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ChannelNumber,
		       PrimaryImageURL, BackdropImageURL
		FROM Items
		WHERE MediaType = 'Channel' AND LocationType = 'Remote'
		  AND (PrimaryImageURL IS NOT NULL OR BackdropImageURL IS NOT NULL)
		ORDER BY ChannelNumber, Name`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var channels []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, primaryImage, backdropImage sql.NullString
		var channelNumber sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &channelNumber,
			&primaryImage, &backdropImage); err != nil {
			return nil, err
		}
		
		channel := map[string]interface{}{
			"Id":            id.String,
			"Name":          name.String,
			"Overview":      overview.String,
			"MediaType":     mediaType.String,
			"ChannelNumber": channelNumber.Int64,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			channel["PrimaryImageUrl"] = primaryImage.String
		}
		if backdropImage.Valid && backdropImage.String != "" {
			channel["BackdropImageUrl"] = backdropImage.String
		}
		
		channels = append(channels, channel)
	}
	
	if channels == nil {
		channels = []map[string]interface{}{}
	}
	
	return channels, rows.Err()
}

// GetProgramsWithImage returns TV programs with images.
func (r *ItemRepository) GetProgramsWithImage(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, RunTimeTicks, PrimaryImageURL
		FROM Items
		WHERE MediaType = 'Video' AND ContentType = 'TvChannel'
		  AND PrimaryImageURL IS NOT NULL
		ORDER BY StartTimeTicks`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var programs []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType, primaryImage sql.NullString
		var startTimeTicks, runTimeTicks sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&startTimeTicks, &runTimeTicks, &primaryImage); err != nil {
			return nil, err
		}
		
		program := map[string]interface{}{
			"Id":              id.String,
			"Name":            name.String,
			"Overview":        overview.String,
			"MediaType":       mediaType.String,
			"Type":            contentType.String,
			"StartTimeTicks":  startTimeTicks.Int64,
			"RunTimeTicks":    runTimeTicks.Int64,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			program["PrimaryImageUrl"] = primaryImage.String
		}
		
		programs = append(programs, program)
	}
	
	if programs == nil {
		programs = []map[string]interface{}{}
	}
	
	return programs, rows.Err()
}

// GetRecordingFolders returns recording folders.
func (r *ItemRepository) GetRecordingFolders(userId string) ([]map[string]interface{}, error) {
	// Return empty for now - recording folders would be a separate concept
	return []map[string]interface{}{}, nil
}

// GetRecommendedPrograms returns recommended TV programs.
func (r *ItemRepository) GetRecommendedPrograms(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       StartTimeTicks, RunTimeTicks, CommunityRating
		FROM Items
		WHERE MediaType = 'Video' AND ContentType = 'TvChannel'
		  AND CommunityRating >= 7.0
		ORDER BY CommunityRating DESC, StartTimeTicks
		LIMIT 20`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var programs []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType sql.NullString
		var startTimeTicks, runTimeTicks sql.NullInt64
		var communityRating sql.NullFloat64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&startTimeTicks, &runTimeTicks, &communityRating); err != nil {
			return nil, err
		}
		
		program := map[string]interface{}{
			"Id":              id.String,
			"Name":            name.String,
			"Overview":        overview.String,
			"MediaType":       mediaType.String,
			"Type":            contentType.String,
			"StartTimeTicks":  startTimeTicks.Int64,
			"RunTimeTicks":    runTimeTicks.Int64,
			"CommunityRating": communityRating.Float64,
		}
		
		programs = append(programs, program)
	}
	
	if programs == nil {
		programs = []map[string]interface{}{}
	}
	
	return programs, rows.Err()
}

// GetPlaylists returns playlists.
func (r *ItemRepository) GetPlaylists(userId string) ([]map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ExtraType,
		       PrimaryImageURL, CreatedDate
		FROM Items
		WHERE MediaType = 'Playlist'
		ORDER BY Name`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var playlists []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, playlistType, primaryImage, createdDate sql.NullString
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &playlistType,
			&primaryImage, &createdDate); err != nil {
			return nil, err
		}
		
		playlist := map[string]interface{}{
			"Id":          id.String,
			"Name":        name.String,
			"Overview":    overview.String,
			"MediaType":   mediaType.String,
			"PlaylistType": playlistType.String,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			playlist["PrimaryImageUrl"] = primaryImage.String
		}
		if createdDate.Valid {
			playlist["CreateDate"] = createdDate.String
		}
		
		playlists = append(playlists, playlist)
	}
	
	if playlists == nil {
		playlists = []map[string]interface{}{}
	}
	
	return playlists, rows.Err()
}

// GetPlaylist returns a single playlist.
func (r *ItemRepository) GetPlaylist(id string) (map[string]interface{}, error) {
	query := `
		SELECT Id, Name, Overview, MediaType, ExtraType as PlaylistType,
		       PrimaryImageURL, BackdropImageURL, CreatedDate, ModifiedDate
		FROM Items
		WHERE Id = ? AND MediaType = 'Playlist'`

	var dbId, name, overview, mediaType, playlistType, primaryImage, backdropImage, createdDate, modifiedDate sql.NullString

	err := r.QueryRow(query, id).Scan(&dbId, &name, &overview, &mediaType, &playlistType,
		&primaryImage, &backdropImage, &createdDate, &modifiedDate)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("playlist not found: %s", id)
	}
	if err != nil {
		return nil, err
	}

	playlist := map[string]interface{}{
		"Id":            dbId.String,
		"Name":          name.String,
		"Overview":      overview.String,
		"MediaType":     mediaType.String,
		"PlaylistType":  playlistType.String,
	}
	
	if primaryImage.Valid && primaryImage.String != "" {
		playlist["PrimaryImageUrl"] = primaryImage.String
	}
	if backdropImage.Valid && backdropImage.String != "" {
		playlist["BackdropImageUrl"] = backdropImage.String
	}
	if createdDate.Valid {
		playlist["CreateDate"] = createdDate.String
	}
	if modifiedDate.Valid {
		playlist["DateModified"] = modifiedDate.String
	}
	
	return playlist, nil
}

// GetPlaylistItems returns items in a playlist.
func (r *ItemRepository) GetPlaylistItems(playlistId string) ([]map[string]interface{}, error) {
	// Playlist items would be in a separate linking table
	// For now, return items that have this playlist as parent
	query := `
		SELECT Id, Name, Overview, MediaType, ContentType,
		       RunTimeTicks, PrimaryImageURL
		FROM Items
		WHERE ParentID = ?
		ORDER BY SortName, Name`
	
	rows, err := r.Query(query, playlistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var items []map[string]interface{}
	for rows.Next() {
		var id, name, overview, mediaType, contentType, primaryImage sql.NullString
		var runTimeTicks sql.NullInt64
		
		if err := rows.Scan(&id, &name, &overview, &mediaType, &contentType,
			&runTimeTicks, &primaryImage); err != nil {
			return nil, err
		}
		
		item := map[string]interface{}{
			"Id":            id.String,
			"Name":          name.String,
			"Overview":      overview.String,
			"MediaType":     mediaType.String,
			"Type":          contentType.String,
			"RunTimeTicks":  runTimeTicks.Int64,
		}
		
		if primaryImage.Valid && primaryImage.String != "" {
			item["PrimaryImageUrl"] = primaryImage.String
		}
		
		items = append(items, item)
	}
	
	if items == nil {
		items = []map[string]interface{}{}
	}
	
	return items, rows.Err()
}

// GetDisplayPreferences returns display preferences.
func (r *ItemRepository) GetDisplayPreferences(userId string) (map[string]interface{}, error) {
	query := `
		SELECT Preferences FROM DisplayPreferences
		WHERE UserId = ?`
	
	var preferences sql.NullString
	
	err := r.QueryRow(query, userId).Scan(&preferences)
	if err == sql.ErrNoRows {
		return map[string]interface{}{
			"UserId":     userId,
			"ScrollDirection": "vertical",
			"ShowBackdrop": true,
			"ShowSidebar": true,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"UserId":      userId,
		"Preferences": preferences.String,
	}, nil
}

// GetDisplayPreferencesByItem returns display preferences for an item.
func (r *ItemRepository) GetDisplayPreferencesByItem(itemId, userId string) (map[string]interface{}, error) {
	query := `
		SELECT Preferences FROM DisplayPreferences
		WHERE UserId = ? AND ItemId = ?`
	
	var preferences sql.NullString
	
	err := r.QueryRow(query, userId, itemId).Scan(&preferences)
	if err == sql.ErrNoRows {
		return map[string]interface{}{
			"UserId": userId,
			"ItemId": itemId,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"UserId":      userId,
		"ItemId":     itemId,
		"Preferences": preferences.String,
	}, nil
}

// GetViewSettings returns view settings.
func (r *ItemRepository) GetViewSettings(userId string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"UserId": userId,
		"Views":  []map[string]interface{}{},
	}, nil
}

// GetGenres returns available genres.
func (r *ItemRepository) GetGenres() ([]map[string]interface{}, error) {
	query := `
		SELECT DISTINCT value FROM (
			SELECT trim(value) as value FROM Items,
			xml_table('genres' passing genres columns value text)
			WHERE genres IS NOT NULL AND genres != ''
		)
		ORDER BY value`
	
	rows, err := r.Query(query)
	if err != nil {
		// Fallback: genres stored as comma-separated in Genres column
		rows, err = r.Query(`
			SELECT DISTINCT Genres FROM Items 
			WHERE Genres IS NOT NULL AND Genres != ''
		`)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()
	
	var genres []map[string]interface{}
	for rows.Next() {
		var genre sql.NullString
		if err := rows.Scan(&genre); err != nil {
			return nil, err
		}
		if genre.Valid {
			genres = append(genres, map[string]interface{}{
				"Name": genre.String,
			})
		}
	}
	
	if genres == nil {
		genres = []map[string]interface{}{}
	}
	
	return genres, rows.Err()
}

// GetStudios returns available studios.
func (r *ItemRepository) GetStudios() ([]map[string]interface{}, error) {
	query := `
		SELECT DISTINCT Studios FROM Items 
		WHERE Studios IS NOT NULL AND Studios != ''
		ORDER BY Studios`
	
	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var studios []map[string]interface{}
	for rows.Next() {
		var studio sql.NullString
		if err := rows.Scan(&studio); err != nil {
			return nil, err
		}
		if studio.Valid {
			studios = append(studios, map[string]interface{}{
				"Name": studio.String,
			})
		}
	}
	
	if studios == nil {
		studios = []map[string]interface{}{}
	}
	
	return studios, rows.Err()
}

// GetYears returns available years.
func (r *ItemRepository) GetYears() ([]int, error) {
	return []int{}, nil
}

// GetNetworks returns available networks.
func (r *ItemRepository) GetNetworks() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetTags returns available tags.
func (r *ItemRepository) GetTags() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetGames returns games.
func (r *ItemRepository) GetGames(userId string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetGame returns a single game.
func (r *ItemRepository) GetGame(id string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

// GetItemMediaTypes returns all media types for an item.
func (r *ItemRepository) GetItemMediaTypes(itemId string) ([]string, error) {
	var mediaTypes []string
	rows, err := r.Query("SELECT MediaType FROM ItemMediaTypes WHERE ItemId = ?", itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mt string
		if err := rows.Scan(&mt); err != nil {
			return nil, err
		}
		mediaTypes = append(mediaTypes, mt)
	}
	if mediaTypes == nil {
		mediaTypes = []string{}
	}
	return mediaTypes, rows.Err()
}

// SetItemMediaTypes sets the media types for an item, replacing any existing types.
func (r *ItemRepository) SetItemMediaTypes(itemId string, mediaTypes []string) error {
	if _, err := r.Exec("DELETE FROM ItemMediaTypes WHERE ItemId = ?", itemId); err != nil {
		return err
	}
	for _, mt := range mediaTypes {
		if _, err := r.Exec("INSERT INTO ItemMediaTypes (ItemId, MediaType) VALUES (?, ?)", itemId, mt); err != nil {
			return err
		}
	}
	return nil
}

// AddItemMediaType adds a single media type to an item.
func (r *ItemRepository) AddItemMediaType(itemId, mediaType string) error {
	_, err := r.Exec("INSERT OR IGNORE INTO ItemMediaTypes (ItemId, MediaType) VALUES (?, ?)", itemId, mediaType)
	return err
}

// GetGameGenres returns game genres.
func (r *ItemRepository) GetGameGenres() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetGameStudios returns game studios.
func (r *ItemRepository) GetGameStudios() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetGameCompanies returns game companies.
func (r *ItemRepository) GetGameCompanies() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// DeleteRecording deletes a recording by ID.
func (r *ItemRepository) DeleteRecording(id string) error {
	return nil
}

// GetTimer returns a timer by ID.
func (r *ItemRepository) GetTimer(id string) (map[string]interface{}, error) {
	return nil, nil
}

// CreateTimer creates a new timer.
func (r *ItemRepository) CreateTimer(req *struct {
	ChannelID    string `json:"ChannelId"`
	ProgramID   string `json:"ProgramId"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	PrePadding  int    `json:"PrePadding"`
	PostPadding int    `json:"PostPadding"`
	Name        string `json:"Name"`
}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"Id":         "timer-1",
		"ChannelId":   req.ChannelID,
		"ProgramId":   req.ProgramID,
		"StartDate":   req.StartDate,
		"EndDate":     req.EndDate,
		"PrePadding":  req.PrePadding,
		"PostPadding": req.PostPadding,
		"Name":        req.Name,
	}, nil
}

// UpdateTimer updates an existing timer.
func (r *ItemRepository) UpdateTimer(id string, req *struct {
	ChannelID    string `json:"ChannelId"`
	ProgramID    string `json:"ProgramId"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	PrePadding  int    `json:"PrePadding"`
	PostPadding int    `json:"PostPadding"`
	Name        string `json:"Name"`
}) error {
	return nil
}

// DeleteTimer deletes a timer by ID.
func (r *ItemRepository) DeleteTimer(id string) error {
	return nil
}

// GetSeriesTimer returns a series timer by ID.
func (r *ItemRepository) GetSeriesTimer(id string) (map[string]interface{}, error) {
	return nil, nil
}

// CreateSeriesTimer creates a new series timer.
func (r *ItemRepository) CreateSeriesTimer(req *struct {
	ChannelID        string `json:"ChannelId"`
	ProgramName     string `json:"ProgramName"`
	StartDate       string `json:"StartDate"`
	EndDate         string `json:"EndDate"`
	PrePadding      int    `json:"PrePadding"`
	PostPadding     int    `json:"PostPadding"`
	Days            []int  `json:"Days"`
	RecordAnyTime   bool   `json:"RecordAnyTime"`
	RecordAnyChannel bool  `json:"RecordAnyChannel"`
}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"Id":                "series-timer-1",
		"ChannelId":        req.ChannelID,
		"ProgramName":      req.ProgramName,
		"StartDate":        req.StartDate,
		"EndDate":          req.EndDate,
		"PrePadding":       req.PrePadding,
		"PostPadding":      req.PostPadding,
		"Days":             req.Days,
		"RecordAnyTime":    req.RecordAnyTime,
		"RecordAnyChannel": req.RecordAnyChannel,
	}, nil
}

// UpdateSeriesTimer updates an existing series timer.
func (r *ItemRepository) UpdateSeriesTimer(id string, req *struct {
	ChannelID        string `json:"ChannelId"`
	ProgramName     string `json:"ProgramName"`
	StartDate       string `json:"StartDate"`
	EndDate         string `json:"EndDate"`
	PrePadding      int    `json:"PrePadding"`
	PostPadding     int    `json:"PostPadding"`
	Days            []int  `json:"Days"`
	RecordAnyTime   bool   `json:"RecordAnyTime"`
	RecordAnyChannel bool  `json:"RecordAnyChannel"`
}) error {
	return nil
}

// DeleteSeriesTimer deletes a series timer by ID.
func (r *ItemRepository) DeleteSeriesTimer(id string) error {
	return nil
}

// GetRecordingSeries returns series recordings.
func (r *ItemRepository) GetRecordingSeries() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetRecordingGroups returns recording groups.
func (r *ItemRepository) GetRecordingGroups() ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// GetRecordingGroup returns a recording group by ID.
func (r *ItemRepository) GetRecordingGroup(id string) (map[string]interface{}, error) {
	return nil, nil
}

// DeleteTunerHost deletes a tuner host by ID.
func (r *ItemRepository) DeleteTunerHost(id string) error {
	return nil
}

// CreateTunerHost creates a new tuner host.
func (r *ItemRepository) CreateTunerHost(req *struct {
	Type    string `json:"Type"`
	Host    string `json:"Host"`
	Port    int    `json:"Port"`
	TunerIP string `json:"TunerIp"`
	Friendly string `json:"FriendlyName"`
}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"Id":          "tuner-1",
		"Type":        req.Type,
		"Host":        req.Host,
		"Port":        req.Port,
		"TunerIp":     req.TunerIP,
		"FriendlyName": req.Friendly,
	}, nil
}

// DeleteListingProvider deletes a listing provider by ID.
func (r *ItemRepository) DeleteListingProvider(id string) error {
	return nil
}

// CreateListingProvider creates a new listing provider.
func (r *ItemRepository) CreateListingProvider(req *struct {
	Type     string `json:"Type"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Country  string `json:"Country"`
	ZipCode  string `json:"ZipCode"`
}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"Id":       "provider-1",
		"Type":     req.Type,
		"Username": req.Username,
		"Country":  req.Country,
		"ZipCode":  req.ZipCode,
	}, nil
}

// CreateChannelMapping creates a channel mapping.
func (r *ItemRepository) CreateChannelMapping(req *struct {
	TunerChannelNumber      string `json:"TunerChannelNumber"`
	ProviderChannelNumber  string `json:"ProviderChannelNumber"`
	ProviderId            string `json:"ProviderId"`
}) error {
	return nil
}
