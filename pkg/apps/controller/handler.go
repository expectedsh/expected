package controller

import (
	"bufio"
	"context"
	"errors"
	"github.com/expectedsh/expected/pkg/apps/controller/docker"
	"github.com/expectedsh/expected/pkg/models/containers"
	"github.com/expectedsh/expected/pkg/protocol"
	"github.com/sirupsen/logrus"
)

func (App) ChangeContainerState(ctx context.Context, r *protocol.ChangeContainerStateRequest) (*protocol.ChangeContainerStateReply, error) {
	container, err := containers.FindContainerByID(ctx, r.Id)
	log := logrus.WithField("id", r.Id).WithField("request", r.RequestedState.String())
	log.Info("new container change request received")

	if err == nil && container == nil {
		err = errors.New("container not found")
	}
	if err != nil {
		log.WithError(err).Error("failed to find container")
		return nil, err
	}

	service, err := docker.ServiceFindByName(ctx, container.ID)
	if err != nil {
		log.WithError(err).Error("failed to find container service")
		return nil, err
	}

	if (service == nil && r.RequestedState == protocol.ChangeContainerStateRequest_STOP) ||
		(service != nil && r.RequestedState == protocol.ChangeContainerStateRequest_START) {
		log.Info("service current state is already to desired state")
		return &protocol.ChangeContainerStateReply{}, nil
	}

	if r.RequestedState == protocol.ChangeContainerStateRequest_START {
		log.Info("creating the service")
		if err := docker.ServiceCreate(ctx, container); err != nil {
			log.WithError(err).Error("failed to create container service")
			return nil, err
		}
	}

	if r.RequestedState == protocol.ChangeContainerStateRequest_STOP {
		log.Info("removing the service")
		if err := docker.ServiceRemove(ctx, container); err != nil {
			log.WithError(err).Error("failed to remove container service")
			return nil, err
		}
	}

	return &protocol.ChangeContainerStateReply{}, nil
}

func (App) GetContainerLogs(r *protocol.GetContainersLogsRequest, ctrl protocol.Controller_GetContainerLogsServer) error {
	container, err := containers.FindContainerByID(ctrl.Context(), r.Id)
	log := logrus.WithField("id", r.Id)
	log.Info("new container logs request received")

	if err == nil && container == nil {
		err = errors.New("container not found")
	}
	if err != nil {
		log.WithError(err).Error("failed to find container")
		return err
	}

	service, err := docker.ServiceFindByName(ctrl.Context(), container.ID)
	if err != nil {
		log.WithError(err).Error("failed to find container service")
		return err
	}
	if service == nil {
		return nil
	}

	logs, err := docker.ServiceGetLogs(ctrl.Context(), service.ID)
	if err != nil {
		log.WithError(err).Error("failed to get container logs")
		return err
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)
	for scanner.Scan() {
		b := scanner.Bytes()
		entry, err := docker.ParseLogEntry(b)
		if err != nil {
			return err
		}
		err = ctrl.Send(&protocol.GetContainersLogsReply{
			Output: entry.Output,
			Message:
		})
		if err != nil {
			return err
		}
	}

	return nil
}
