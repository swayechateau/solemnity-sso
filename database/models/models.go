package models

import "github.com/google/uuid"

func NewUUID() []byte {
	newUUID := uuid.New()
	return newUUID[:]
}

func NewUUIDString() string {
	newUUID := uuid.New()
	return newUUID.String()
}

func UUIDStringToBytes(uuidString string) []byte {
	uuid, _ := uuid.Parse(uuidString)
	return uuid[:]
}

func BytesToUUID(uuidByteSlice []byte) uuid.UUID {
	uuid, _ := uuid.FromBytes(uuidByteSlice)
	return uuid
}
