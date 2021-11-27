//go:build !darwin
// +build !darwin

package matrix

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
			return NewRGBLedMatrix(config)
		}
	}
}
