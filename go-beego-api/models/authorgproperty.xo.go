// Package models contains the types for schema 'dbo'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Authorgproperty represents a row from 'dbo.AuthOrgProperty'.
type Authorgproperty struct {
	ID          string    `json:"Id"`          // Id
	Orgid       string    `json:"OrgId"`       // OrgId
	Name        string    `json:"Name"`        // Name
	Value       string    `json:"Value"`       // Value
	Revision    int       `json:"Revision"`    // Revision
	Createdby   string    `json:"CreatedBy"`   // CreatedBy
	Createdtime time.Time `json:"CreatedTime"` // CreatedTime
	Updatedby   string    `json:"UpdatedBy"`   // UpdatedBy
	Updatedtime time.Time `json:"UpdatedTime"` // UpdatedTime

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Authorgproperty exists in the database.
func (a *Authorgproperty) Exists() bool {
	return a._exists
}

// Deleted provides information if the Authorgproperty has been deleted from the database.
func (a *Authorgproperty) Deleted() bool {
	return a._deleted
}

// Insert inserts the Authorgproperty to the database.
func (a *Authorgproperty) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO dbo.AuthOrgProperty (` +
		`Id, OrgId, Name, Value, Revision, CreatedBy, CreatedTime, UpdatedBy, UpdatedTime` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)`

	// run query
	XOLog(sqlstr, a.ID, a.Orgid, a.Name, a.Value, a.Revision, a.Createdby, a.Createdtime, a.Updatedby, a.Updatedtime)
	_, err = db.Exec(sqlstr, a.ID, a.Orgid, a.Name, a.Value, a.Revision, a.Createdby, a.Createdtime, a.Updatedby, a.Updatedtime)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Authorgproperty in the database.
func (a *Authorgproperty) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE dbo.AuthOrgProperty SET ` +
		`OrgId = $1, Name = $2, Value = $3, Revision = $4, CreatedBy = $5, CreatedTime = $6, UpdatedBy = $7, UpdatedTime = $8` +
		` WHERE Id = $9`

	// run query
	XOLog(sqlstr, a.Orgid, a.Name, a.Value, a.Revision, a.Createdby, a.Createdtime, a.Updatedby, a.Updatedtime, a.ID)
	_, err = db.Exec(sqlstr, a.Orgid, a.Name, a.Value, a.Revision, a.Createdby, a.Createdtime, a.Updatedby, a.Updatedtime, a.ID)
	return err
}

// Save saves the Authorgproperty to the database.
func (a *Authorgproperty) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Delete deletes the Authorgproperty from the database.
func (a *Authorgproperty) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM dbo.AuthOrgProperty WHERE Id = $1`

	// run query
	XOLog(sqlstr, a.ID)
	_, err = db.Exec(sqlstr, a.ID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AuthorgpropertyByID retrieves a row from 'dbo.AuthOrgProperty' as a Authorgproperty.
//
// Generated from index 'PK_AuthOrgProperty'.
func AuthorgpropertyByID(db XODB, id string) (*Authorgproperty, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`Id, OrgId, Name, Value, Revision, CreatedBy, CreatedTime, UpdatedBy, UpdatedTime ` +
		`FROM dbo.AuthOrgProperty ` +
		`WHERE Id = $1`

	// run query
	XOLog(sqlstr, id)
	a := Authorgproperty{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &a.Orgid, &a.Name, &a.Value, &a.Revision, &a.Createdby, &a.Createdtime, &a.Updatedby, &a.Updatedtime)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
