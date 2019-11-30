package mongo

import "github.com/jpw547/indecision/structs"

// CreateHistoryRecord makes a record of a choice in the database
func (m *DB) CreateHistoryRecord(record structs.HistoryRecord) (structs.HistoryRecord, error) {
	var toReturn structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// GetHistoryRecord returns a record from the database
func (m *DB) GetHistoryRecord(id string) (structs.HistoryRecord, error) {
	var toReturn structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// UpdateHistoryRecord updates a record in the database
func (m *DB) UpdateHistoryRecord(id string, record structs.HistoryRecord) (structs.HistoryRecord, error) {
	var toReturn structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// DeleteHistoryRecord deletes a record from the database
func (m *DB) DeleteHistoryRecord(id string) error {
	return nil
}

// GetHistoryRecordsByUser gets all of the records associated with a certain user
func (m *DB) GetHistoryRecordsByUser(userID string) ([]structs.HistoryRecord, error) {
	var toReturn []structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// GetLastRecordByUser gets the last choice recorded for that user
func (m *DB) GetLastRecordByUser(userID string) (structs.HistoryRecord, error) {
	var toReturn structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// GetHistoryRecordsByChoice gets all of the records associated with a certain choice
func (m *DB) GetHistoryRecordsByChoice(choiceID string) ([]structs.HistoryRecord, error) {
	var toReturn []structs.HistoryRecord

	// do the thing

	return toReturn, nil
}

// GetLastRecordByChoice gets the last record in the database for a when this choice was chosen by a user
func (m *DB) GetLastRecordByChoice(choiceID string) (structs.HistoryRecord, error) {
	var toReturn structs.HistoryRecord

	// do the thing

	return toReturn, nil
}
