package server

import (
	"errors"
	"fmt"
	"net"
)

var (
	ErrPortInUse = errors.New("error: Port already in use")
)


// Parsing

func (server Server) Parse() (Server,error) {
	err := server.Validate()
	if err != nil {
		return server,err
	}
	server.ApiServer,err = server.ApiServer.Parse()
	if err != nil {
		return server,err
	}
	server.WebServer,err = server.WebServer.Parse()
	if err != nil {
		return server,err
	}
	if server.IsPortInUse() {
		if server.ForcePort {
			return server,ErrPortInUse
		} else {
			server.Port = server.GetNextAvailablePort()
		}
	}
	return server,nil
} 

func (apiServer ApiServer) Parse() (ApiServer,error) {
	return apiServer,nil
} 


// Validation

func (server Server) Validate() (error) {
	err := server.WebServer.Validate()
	if err != nil {
		return err
	}
	err = server.ApiServer.Validate()
	if err != nil {
		return err
	}
	if server.IsPortInUse() && server.ForcePort {
		return ErrPortInUse
	}
	return nil	
}

func (apiServer ApiServer) Validate() (error) {
	return nil	
}


// Port Validation

func (server Server) IsPortInUse() bool {
	// If the listener fails to start the port is already in use
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip,server.Port))
	if err != nil {
		return true
	}
	listener.Close()
	return false
}

func (server Server) GetNextAvailablePort() uint16 {
	for {
		if !server.IsPortInUse() {
			return server.Port
		}
		server.Port++
	}
}
