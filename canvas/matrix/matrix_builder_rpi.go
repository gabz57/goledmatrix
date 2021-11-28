//go:build !darwin
// +build !darwin

package matrix

func BuildMatrix(config *MatrixConfig) (Matrix, error) {
	if config.Client {
		matrix, err := NewRpcClient(config)
		if err != nil {
			return nil, err
		}
		if err = validateGeometry(config, matrix); err != nil {
			return nil, err
		}
		return matrix, nil
	} else {
		if config.Emulator == true {
			return NewEmulator(config)
		} else {
			return NewRGBLedMatrix(config)
		}
	}
}
