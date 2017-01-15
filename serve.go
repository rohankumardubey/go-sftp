package sftp

import (
	"io"

	"github.com/nethack42/go-sftp/sshfxp"
)

func readConn(r io.Reader, ch chan<- sshfxp.Packet) error {
	for {
		var pkt sshfxp.Packet

		if err := pkt.Read(r); err != nil {
			return err
		}

		ch <- pkt
	}
	return nil
}

func writeConn(w io.Writer, ch <-chan sshfxp.Packet) error {
	for pkt := range ch {
		blob, err := pkt.Bytes()
		if err != nil {
			return err
		}

		if _, err := w.Write(blob); err != nil {
			return err
		}
	}
	return nil
}