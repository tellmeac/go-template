package entities

import "github.com/google/uuid"

type User struct {
	UID  uuid.UUID
	Name string
}
