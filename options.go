package tlocales

func UseDriver(driver ReaderDriver) Option {
	return func(self *loadContract) {
		self.reader = driver
	}
}

func RegisterUnmarshaller(extension string, umr Unmarshaller) Option {
	return func(self *loadContract) {
		self.unmarshallers[extension] = umr
	}
}
