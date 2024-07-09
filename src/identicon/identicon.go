package identicon

func GenerateIdenticon(accountId string) error {
	hash := generateMd5Hash(accountId)

	_ = hash.getForeground()
	_, err := hash.getColor()
	if err != nil {
		return err
	}

	return nil
}
