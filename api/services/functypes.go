package services

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
type InfoFunc func(UserServiceDataToCheck) (AccountInfo, error)
