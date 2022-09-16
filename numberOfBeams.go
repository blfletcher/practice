package main

// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/

// algorithm
/***********

Initialize a previousDeviceCounter value to 0
Initialize laserbeams to 0
Iterate through the string array
If it contains 1s:
  - count 1s, multiply by previousDeviceCounter and add result to laserbeams
  - set previousDeviceCounter to current 1 count

return laserbeams

************/

func numberOfBeams(bank []string) int {
	laserBeams := 0
	prevDevices := 0

	for _, v := range bank {
		val := v

		devices := 0
		for _, j := range val {
			jval := j

			if jval == '1' {
				devices++
			}
		}

		if devices > 0 {
			laserBeams += prevDevices * devices
			prevDevices = devices
		}
	}

	return laserBeams
}
