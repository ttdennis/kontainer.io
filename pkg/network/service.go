// Package network handles container networks and interconnections
package network

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	networkTypes "github.com/docker/docker/api/types/network"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
)

var (
	// ErrNetworkNotExist occurs when a network does not exist
	ErrNetworkNotExist = errors.New("Network does not exist")

	// ErrNetworkAlreadyExists occurs when a network already exists
	ErrNetworkAlreadyExists = errors.New("Network already exists")
)

// Service NetworkService
type Service interface {
	// CreateNetwork creates a new network for a given user
	CreateNetwork(refid uint, cfg *Config) error

	// RemoveNetwork removes a network with a given name
	RemoveNetworkByName(refid uint, name string) error

	// AddContainerToNetwork joins a given container to a given network
	AddContainerToNetwork(refid uint, name string, containerID string) error

	// RemoveContainerFromNetwork removes a container from a given network
	RemoveContainerFromNetwork(refid uint, name string, containerID string) error

	// ExposePortToContainer exposes a port from one container to another
	ExposePortToContainer(refid uint, srcContainerID string, port uint32, destContainerID string) error

	// RemovePortFromContainer removes an exposed port from a container
	RemovePortFromContainer(refid uint, srcContainerID string, port uint32, destContainerID string) error

	// UserHasNetwork checks whether a user has created a network
	UserHasNetwork(refid uint) bool
}

type dbAdapter interface {
	abstraction.DBAdapter
	AutoMigrate(...interface{}) error
	Where(interface{}, ...interface{}) error
	First(interface{}, ...interface{}) error
	Create(interface{}) error
	Delete(interface{}, ...interface{}) error
}

type service struct {
	dcli   abstraction.DCli
	db     dbAdapter
	logger log.Logger
}

func (s *service) InitializeDatabases() error {
	return s.db.AutoMigrate(&Networks{})
}

func (s *service) getNetworkByName(refid uint, name string) (Networks, error) {
	nw := Networks{}

	err := s.db.Where("UserID = ? AND NetworkName = ?", refid, name)
	if err != nil {
		return nw, err
	}

	s.db.First(&nw)

	return nw, nil
}

func (s *service) CreateNetwork(refid uint, cfg *Config) error {
	name := cfg.Name

	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	if nw.NetworkID != "" {
		return ErrNetworkAlreadyExists
	}

	res, err := s.dcli.NetworkCreate(context.Background(), fmt.Sprintf("%s-%s", string(refid), name), types.NetworkCreate{
		Driver: cfg.Driver,
	})
	if err != nil {
		return err
	}

	nw = Networks{
		UserID:      uint(refid),
		NetworkName: name,
		NetworkID:   res.ID,
	}

	err = s.db.Create(&nw)
	if err != nil {
		// Try to remove the actual network on db error
		s.dcli.NetworkRemove(context.Background(), res.ID)
		return err
	}

	return nil
}

func (s *service) RemoveNetworkByName(refid uint, name string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	if nw.NetworkID != "" {
		err = s.dcli.NetworkRemove(context.Background(), nw.NetworkID)
		if err != nil {
			return err
		}

		err = s.db.Delete(&nw)
		if err != nil {
			return err
		}

		return nil
	}

	return ErrNetworkNotExist
}

func (s *service) AddContainerToNetwork(refid uint, name string, containerID string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}

	if nw.NetworkID != "" {
		err = s.dcli.NetworkConnect(context.Background(), nw.NetworkID, containerID, &networkTypes.EndpointSettings{})
		if err != nil {
			return err
		}
	} else {
		return ErrNetworkNotExist
	}

	return nil
}

func (s *service) RemoveContainerFromNetwork(refid uint, name string, containerID string) error {
	nw, err := s.getNetworkByName(refid, name)
	if err != nil {
		return err
	}
	if nw.NetworkID != "" {
		err = s.dcli.NetworkDisconnect(context.Background(), nw.NetworkID, containerID, true)
		if err != nil {
			return err
		}
	} else {
		return ErrNetworkNotExist
	}

	return nil
}

func (s *service) ExposePortToContainer(refid uint, srcContainerID string, port uint32, destContainerID string) error {
	// TODO: implement
	return nil
}

func (s *service) RemovePortFromContainer(refid uint, srcContainerID string, port uint32, destContainerID string) error {
	// TODO: implement
	return nil
}

func (s *service) UserHasNetwork(refid uint) bool {
	nw := Networks{}

	err := s.db.Where("UserID = ?", refid)
	if err != nil {
		return false
	}

	err = s.db.First(&nw)
	if err != nil {
		return false
	}

	return true
}

// NewService creates a new network service
func NewService(dcli abstraction.DCli, db dbAdapter) (Service, error) {
	s := &service{
		dcli: dcli,
		db:   db,
	}

	err := s.InitializeDatabases()
	if err != nil {
		return s, err
	}

	return s, nil
}
