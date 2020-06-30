/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2020/06/13 18:29
*/

package timestamp

import "time"

type Timestamp struct {
	CreatedAt time.Time `json:"created_at,omitempty"bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"bson:"deleted_at,omitempty"`
}

func Create() Timestamp {
	return Timestamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Time{},
	}
}

func Update(timestamp Timestamp) Timestamp {
	timestamp.UpdatedAt = time.Now()
	return timestamp
}

func Delete(timestamp Timestamp) Timestamp {
	timestamp.DeletedAt = time.Now()
	return timestamp
}
func Restore(timestamp Timestamp) Timestamp {
	timestamp.DeletedAt = time.Time{}
	return timestamp
}
