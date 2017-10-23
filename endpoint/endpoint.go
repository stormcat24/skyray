package endpoint

import (
	"log"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
	"github.com/stormcat24/skyray/iohelper"
	proto "github.com/stormcat24/skyray/proto/skyray"
)

type SkyrayEndpoint struct {
}

func NewSkyrayEndpoint() proto.SkyrayServiceServer {
	return &SkyrayEndpoint{}
}

func (e SkyrayEndpoint) Connect(req *proto.Command, stream proto.SkyrayService_ConnectServer) error {

	callback := func(data []byte) {
		res := proto.Response{
			Output: data,
		}
		if err := stream.Send(&res); err != nil {
			log.Println(err.Error())
		}
	}

	stdout := iohelper.NewBufferObserver(callback)
	cmd := exec.CommandContext(stream.Context(), req.Command, req.Arguments...)
	cmd.Stdout = stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	err := cmd.Start()
	if err != nil {
		return errors.Wrapf(err, "Receive command is failed.")
	}

	go func() {
		select {
		case <-stream.Context().Done():
			if err := cmd.Process.Signal(syscall.SIGKILL); err != nil {
				log.Println(err.Error())
			}
			return
		}
	}()

	if err := cmd.Wait(); err != nil {
		return errors.Wrapf(err, "Waid command is failed.")
	}
	return nil
}
