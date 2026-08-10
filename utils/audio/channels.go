package audio

// UpdateChannelNumberTo2 updates the data to two channels.
// This is just a temporary solution for the problem mentioned in https://github.com/braheezy/shine-mp3/issues/11.
func UpdateChannelNumberTo2(
	data []int16,
) []int16 {
	newData := make([]int16, len(data)*2)
	for i, sample := range data {
		newData[i*2] = sample
		newData[i*2+1] = sample
	}
	return newData
}
