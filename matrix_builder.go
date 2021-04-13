package goledmatrix

import (
	"errors"
	"strconv"
)

func BuildMatrix(config *MatrixConfig) (Matrix, error) {
	if config.Client {
		client, err := NewMatrixRpcClient(config)
		if err != nil {
			return nil, err
		}
		err = validateGeometry(config, client)
		if err != nil {
			return nil, err
		}
		return client, nil
	} else {
		if config.Emulator == true {
			return NewMatrixEmulator(config)
		} else {
			return NewMatrixEmulator(config)
		}
	}
}

func validateGeometry(config *MatrixConfig, remoteMatrix Matrix) error {
	width, height := config.Geometry()
	clientW, clientH := remoteMatrix.Geometry()
	if width != clientW {
		return errors.New("incorrect WIDTH detected between local (" + strconv.Itoa(width) + ") and received from remote (" + strconv.Itoa(clientW) + ")")
	}
	if height != clientH {
		return errors.New("incorrect HEIGHT detected between local (" + strconv.Itoa(height) + ") and received from remote (" + strconv.Itoa(clientH) + ")")
	}
	return nil
}
