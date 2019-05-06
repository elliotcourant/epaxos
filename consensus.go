package epaxos

import (
	"github.com/kataras/golog"
	"github.com/readystock/goqu"
	"google.golang.org/grpc"
)

// Cluster is the EPaxos implementation object.
type cluster struct {
	grpcServer grpc.Server

	commandChannel <-chan CommandRequest
}

func (c *cluster) runNode() {
	for {
		select {
		case command := <-c.commandChannel:
			c.performCommand(command)
		}
	}
}

func (c *cluster) performCommand(msg CommandRequest) {
	potentialSQLCommand, sqlOk := msg.GetPotentialSqlCommand().(CommandRequest_SqlCommand)
	if sqlOk {
		sqlCommand := *potentialSQLCommand.SqlCommand
		switch sqlCommand.GetAction() {
		case SqlAction_INSERT:

		case SqlAction_UPDATE:

		case SqlAction_DELETE:

		default:
			golog.Errorf("could not handle action for command")
		}
	}

	potentialSetCommand, setOk := msg.GetPotentialSetCommand().(CommandRequest_SetCommand)
	if setOk {

	}

	potentialDeleteCommand, deleteOk := msg.GetPotentialDeleteCommand().(CommandRequest_DeleteCommand)
	if deleteOk {

	}
}
